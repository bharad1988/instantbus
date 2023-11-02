package main

import (
	"github.com/bharad1988/instantbus/testpack"
)

func main() {
	var wait chan bool
	ss := testpack.SubscriptionService{}
	ss.StartSubscriptionService("localhost:5001")
	ss.SubscribeTopic("world")
	ss.SubscribeTopic("earth")
	<-wait
}
