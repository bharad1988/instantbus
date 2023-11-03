# instantbus
A standalone service bus over network written in golang.

**_caveat: No error handling yet. Data races reported with -race option have been fixed_** 

## Service Bus component
- Allows services to subsribe to a topic .
- Receives messages in byte array over a topic (serialized), pushes it to all subscribers for that topic
- Subscribers need to be aware of de-serialization method.
### How to start bus service
```
cd busservice
go run main.go
```

## Producer/Publisher
Produces a message for a topic

## Consumer/Subsrciber
Subscribes to a topic over the bus via n/w and get messages as they arrive.

## Example usage of SDK
sclient directory/package has the SDK for publisher and subscriber
```
cd example
go run main.go
// server starts at "localhost:5001"
```
### Subscribe to bus service and a topic
```
	ss := sbclient.SubscriptionService{}
 	ss.StartSubscriptionService("localhost:5001")
 	ss.SubscribeTopic("world")
	ss.SubscribeTopic("earth")
```
### Add a publisher to bus service
```
  	ss2 := sbclient.SubscriptionService{}
	ss2.StartPublisher("localhost:5001")
	for i := 0; i < 100; i++ {
		m := fmt.Sprintf("test %d", i)
		testMsg := []byte(m)
        // send message to topic , e.g: "world" / "earth"
		ss2.SendMessage("world", testMsg)
		ss2.SendMessage("earth", testMsg)
	}
```

**More sample code with encrypt and decrypt**

Encrypt a message and send
```
	fmt.Println("This is an original:", plainText)
	encrypted, err := GetAESEncrypted(plainText)
	if err != nil {
		fmt.Println("Error during encryption", err)
	}

	ss2.SendMessage("earth", encrypted)
	fmt.Println("This is an encrypted:", encrypted)
```
Receive all messages
```
	fmt.Println("len of messages ", len(ss.GetAllMessages("earth")))
	for i := 0; i < len(ss.GetAllMessages("earth")); i++ {
		fmt.Printf("All message - %s", ss.GetAllMessages("earth")[i])
	}

```
Receive only unread messages
```
	newMessages := ss.GetAllUnreadMessages("earth")
	for _, m := range newMessages {
		fmt.Printf("topic : earth : message %s \n", m)
	}
```
Receive and decrypt messages
```
	ss2.SendMessage("earth", encrypted)
	ss.GetAllUnreadMessages("earth")
	for _, m := range newMessages {
		d, _ := GetAESDecrypted(m)
		fmt.Printf("topic : earth : message %s \n", d)
	}
```
reference: For encryption and descryption https://medium.com/insiderengineering/aes-encryption-and-decryption-in-golang-php-and-both-with-full-codes-ceb598a34f41
