# instantbus
A standalone service bus over network written in golang

## Service Bus component
- Allows services to subsribe to a topic .
- Receives messages in byte array over a topic (serialized), pushes it to all subscribers for that topic
- Subscribers need to be aware of de-serialization method.
How to start bus service
'''
cd busservice
go run main.go
'''

## Producer/Publisher
Produces a message for a topic

## Consumer/Subsriber
Subscribes to a topic over the bus via n/w and get messages as they arrive.

## Example usage of SDK
`
cd subs
go run main.go
`
### Subscribe to bus service and a topic
`
	ss := sbclient.SubscriptionService{}
 	ss.StartSubscriptionService("localhost:5001")
 	ss.SubscribeTopic("world")
	ss.SubscribeTopic("earth")
`
### Add a publisher to bus service
`
  	ss2 := sbclient.SubscriptionService{}
	ss2.StartPublisher("localhost:5001")
	for i := 0; i < 100; i++ {
		m := fmt.Sprintf("test %d", i)
		testMsg := []byte(m)
        // send message to topic , e.g: "world" / "earth"
		ss2.SendMessage("world", testMsg)
		ss2.SendMessage("earth", testMsg)
	}
`
