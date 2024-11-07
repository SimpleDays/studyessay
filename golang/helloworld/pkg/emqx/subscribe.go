/**
 * @Author: birney.dong
 * @Date: 2024/4/22 09:59
 * @Description: subscribe
 */

package emqx

import (
	"context"
	"fmt"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"net/url"
	"strconv"
	"sync"
	"time"
)

var subLock sync.Mutex

func Subscribe(ctx context.Context, serverURL *url.URL, subMap map[string]paho.SubscribeOptions, clientId int, sub chan *paho.Publish) {
	subLock.Lock()

	defer subLock.Unlock()

	done := make(chan struct{})

	cliCfg := autopaho.ClientConfig{
		BrokerUrls: []*url.URL{serverURL},
		KeepAlive:  20, // Keepalive message should be sent every 20 seconds
		//CleanStartOnInitialConnection: true, // Previous tests should not contaminate this one!
		//SessionExpiryInterval:         60,   // If connection drops we want session to remain live whilst we reconnect
		OnConnectionUp: func(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
			fmt.Println(fmt.Sprintf("mqtt connection ok,subMap is %+v", subMap))
			if _, err := cm.Subscribe(context.Background(), &paho.Subscribe{
				Subscriptions: subMap,
			}); err != nil {
				panic(err)
			}
			fmt.Println("mqtt subscribe ok")
			done <- struct{}{}
		},
		OnConnectError: func(err error) { fmt.Printf("subscribe: error whilst attempting connection: %s\n", err) },
		// Errors:         logger{prefix: "subscribe"},

		// eclipse/paho.golang/paho provides base mqtt functionality, the below config will be passed in for each connection
		ClientConfig: paho.ClientConfig{
			ClientID: "TestSub" + strconv.FormatInt(time.Now().UnixMilli(), 10),
			Router: paho.NewSingleHandlerRouter(func(m *paho.Publish) {
				// time.Sleep(2 * time.Second)
				// fmt.Println(fmt.Sprintf("subscribe: clientId: %v, topic: %s,  mqtt message received, msg: %s %s", clientId, m.Topic, time.Now().Format(time.RFC3339), string(m.Payload)))
				sub <- m
			}),
			OnClientError: func(err error) { fmt.Printf("subscribe: client error: %s\n", err) },
			OnServerDisconnect: func(d *paho.Disconnect) {
				if d.Properties != nil {
					fmt.Printf("subscribe: server requested disconnect: %s\n", d.Properties.ReasonString)
				} else {
					fmt.Printf("subscribe:server requested disconnect; reason code: %d\n", d.ReasonCode)
				}
			},
		},
	}

	c, err := autopaho.NewConnection(ctx, cliCfg)
	if err != nil {
		fmt.Println("subscribe: error creating connection: ", err)
		panic(err)
	}

	fmt.Println("subscribe: connection created, c: ", c)

	<-done
}
