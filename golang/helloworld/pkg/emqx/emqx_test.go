/**
 * @Author: birney.dong
 * @Date: 2024/4/22 10:00
 * @Description: emqx_test.go
 */

package emqx

import (
	"context"
	"fmt"
	"github.com/eclipse/paho.golang/paho"
	"net/url"
	"testing"
	"time"
)

//package main
//
//import (
//	"context"
//	"fmt"
//	"github.com/eclipse/paho.golang/paho"
//	"net/url"
//	"sync"
//	"testgo/emqx"
//)

var addPrefix = []string{"c1", "c2", "c3"}

func TestPublish(t *testing.T) {
	//ctx, stop := signal.NotifyContext( , os.Interrupt, syscall.SIGTERM)

	exit := make(chan struct{})
	ctx := context.Background()

	u, err := url.Parse("mqtt://172.16.18.3:1883")
	if err != nil {
		panic(err)
	}

	//subMap := make([]paho.SubscribeOptions, 0)
	//subMap = append(subMap, paho.SubscribeOptions{
	//	Topic: "$share/g1/birney/dong/+/send",
	//	QoS:   1,
	//})
	subMap := make(map[string]paho.SubscribeOptions)
	subMap["$share/g1/birney/dong/+/send"] = paho.SubscribeOptions{QoS: 1}

	// emqx.Subscribe(ctx, u, subMap, 2)
	//wg := sync.WaitGroup{}
	//
	//wg.Add(2)
	//for i := 1; i < 3; i++ {
	//	go func(c int) {
	//		emqx.Subscribe(ctx, u, subMap, c)
	//		fmt.Println(fmt.Sprintf("clientId: %v, sub success", c))
	//		wg.Done()
	//	}(i)
	//}
	//
	//wg.Wait()

	var subChan = make(chan *paho.Publish)
	Subscribe(ctx, u, subMap, 1, subChan)
	fmt.Println(fmt.Sprintf("clientId: %v, sub success", 1))

	go func() {
		defer close(subChan)

		for {
			select {
			case s := <-subChan:
				fmt.Println(fmt.Sprintf("subscribe: clientId: %v, topic: %s,  mqtt message received, msg: %s %s", 1, s.Topic, time.Now().Format(time.RFC3339), string(s.Payload)))
			}
		}
	}()

	//for i := 0; i < 3; i++ {
	//	go func(s int) {
	//		for {
	//			Publish(ctx, u, fmt.Sprintf("[%s hello%s]", time.Now().Format(time.RFC3339), "world"), fmt.Sprintf("$delayed/25/birney/dong/test%v/send", s+1))
	//
	//			time.Sleep(300 * time.Millisecond)
	//
	//			//fmt.Println(fmt.Sprintf("publish success, i: %d", i))
	//		}
	//	}(i)
	//}

	for {
		Publish(ctx, u, fmt.Sprintf("[%s hello%s]", time.Now().Format(time.RFC3339), "world"), fmt.Sprintf("$delayed/5/birney/dong/test%v/send", 1))

		time.Sleep(1 * time.Second)

		//fmt.Println(fmt.Sprintf("publish success, i: %d", i))
	}

	// emqx.Publish(ctx, u, "hello", "birney/dong/test1/send")

	<-exit
	// fmt.Println("hello")
}
