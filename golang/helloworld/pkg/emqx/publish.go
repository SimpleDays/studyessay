/**
 * @Author: birney.dong
 * @Date: 2024/4/22 09:59
 * @Description: publish.go
 */

package emqx

import (
	"context"
	"fmt"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"helloworld/pkg/emqx/log"
	"net/url"
	"strconv"
	"sync"
	"time"
)

var mqttPublishObj *autopaho.ConnectionManager
var lock sync.Once
var push = make(chan struct{})

func Publish(ctx context.Context, serverURL *url.URL, msg interface{}, topic string) {
	lock.Do(func() {
		fmt.Println("publish: create mqtt connection")

		// defer close(push)
		cliCfg := autopaho.ClientConfig{
			BrokerUrls: []*url.URL{serverURL},
			KeepAlive:  5, // Keepalive message should be sent every 20 seconds
			OnConnectionUp: func(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
				fmt.Println("publish: mqtt connection up")
				push <- struct{}{}
			},
			OnConnectError: func(err error) { fmt.Printf("publish: error whilst attempting connection err: %s\n", err) },
			Debug:          log.NewEMQXDebugLogger("EMQX - Debug"),
			PahoErrors:     log.NewEMQXDebugLogger("EMQX - PahoErrors"),
			PahoDebug:      log.NewEMQXDebugLogger("EMQX - PahoDebug"),
			// eclipse/paho.golang/paho provides base mqtt functionality, the below config will be passed in for each connection
			ClientConfig: paho.ClientConfig{
				ClientID:      "TestPub" + strconv.FormatInt(time.Now().UnixMilli(), 10),
				OnClientError: func(err error) { fmt.Printf("publish: client error: %s\n", err) },
				OnServerDisconnect: func(d *paho.Disconnect) {
					if d.Properties != nil {
						fmt.Printf("publish: server requested disconnect: %s\n", d.Properties.ReasonString)
					} else {
						fmt.Printf("publish:server requested disconnect; reason code: %d\n", d.ReasonCode)
					}
				},
			},
		}

		c, err := autopaho.NewConnection(ctx, cliCfg)
		if err != nil {
			panic(err)
		}

		<-push

		mqttPublishObj = c
	})

	_, err := mqttPublishObj.Publish(ctx, &paho.Publish{
		Topic:   topic,
		QoS:     1,
		Payload: []byte(fmt.Sprintf("%v", msg)),
	})

	if err != nil {
		fmt.Println("publish: error while publishing message, err: ", err)
	}

	// fmt.Println(fmt.Sprintf("publish: message published, p is: %+v", p))

}
