package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/bharad1988/instantbus/busproto"
	sub_pb "github.com/bharad1988/instantbus/subscriber"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 5001, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedBusServer
}

type subStatus struct {
	subscriber string
	index      int
}

var TopicMessage map[string][][]byte
var TopicSubs map[string][]subStatus
var SubConns map[string]sub_pb.PushClient

var messagesLock sync.RWMutex

// SayHello implements helloworld.GreeterServer
func (s *server) SendMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	log.Printf("Received: %v", in.GetTopic())
	log.Printf("Received: %v", in.GetMessage())
	messagesLock.Lock()
	TopicMessage[in.GetTopic()] = append(TopicMessage[in.GetTopic()], in.GetMessage())
	messagesLock.Unlock()
	e := time.Now().Unix()
	return &pb.MessageReply{Ts: &e}, nil
}

func generateSubstatus(topic, subscriber string) {
	TopicSubs[topic] = append(TopicSubs[topic], subStatus{subscriber: subscriber, index: 0})

}

// SayHello implements helloworld.GreeterServer
func (s *server) SubTopic(ctx context.Context, in *pb.SubscribeRequest) (*pb.SubscribeReply, error) {
	log.Printf("Received: %v", in.GetTopic())
	log.Printf("Received: %v", in.GetSubscriber())
	generateSubstatus(in.GetTopic(), in.GetSubscriber())
	status := true
	return &pb.SubscribeReply{Done: &status}, nil
}

func main() {
	flag.Parse()
	TopicSubs = make(map[string][]subStatus)
	TopicMessage = make(map[string][][]byte)
	SubConns = make(map[string]sub_pb.PushClient)
	go pushLoop()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBusServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func pushLoop() {
	for {
		pushMessages()
		time.Sleep(time.Millisecond * 100)
	}
}

func pushMessages() {
	// log.Println("Pushing all messages")
	messagesLock.RLock()
	for topic, messages := range TopicMessage {
		subs := TopicSubs[topic]
		for j, sub := range subs {
			for i := sub.index; i < len(messages); i++ {
				pushMessageToSub(sub.subscriber, topic, messages[i])
				sub.index++
			}
			subs[j] = sub
		}
	}
	messagesLock.RUnlock()
}

func createCon(sub string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(sub, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		c := sub_pb.NewPushClient(conn)
		SubConns[sub] = c
	}
	// defer conn.Close()
}

func pushMessageToSub(sub string, topic string, message []byte) {
	if SubConns[sub] == nil {
		createCon(sub)
	}
	c := SubConns[sub]

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.PushMessage(ctx, &sub_pb.PushMessageRequest{Topic: &topic, Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Pushed: %d", r.GetTs())

}
