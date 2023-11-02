package main

import (
	"fmt"

	"github.com/bharad1988/instantbus/sbclient"
)

func main() {
	var wait chan bool
	ss := sbclient.SubscriptionService{}
	ss2 := sbclient.SubscriptionService{}
	ss2.StartPublisher("localhost:5001")
	ss.StartSubscriptionService("localhost:5001")
	ss.SubscribeTopic("world")
	ss.SubscribeTopic("earth")
	for i := 0; i < 100; i++ {
		m := fmt.Sprintf("test %d", i)
		testMsg := []byte(m)
		ss2.SendMessage("world", testMsg)
		ss2.SendMessage("earth", testMsg)
	}
	<-wait
}
