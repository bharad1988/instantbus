package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	bus_pb "github.com/bharad1988/instantbus/busproto"
	pb "github.com/bharad1988/instantbus/subscriber"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:5001", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
	port int
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedPushServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) PushMessage(ctx context.Context, in *pb.PushMessageRequest) (*pb.PushMessageReply, error) {
	log.Printf("Received: %v", in.GetTopic())
	log.Printf("Received: %v", string(in.GetMessage()))
	e := time.Now().Unix()
	return &pb.PushMessageReply{Ts: &e}, nil
}

func main() {
	flag.Parse()
	rand.NewSource(time.Now().UnixNano())

	// generate random integer between 10 and 20
	port = rand.Intn(60000) + 1024
	subscribeTopic()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPushServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func subscribeTopic() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := bus_pb.NewBusClient(conn)
	clientID := fmt.Sprintf("localhost:%d", port)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SubTopic(ctx, &bus_pb.SubscribeRequest{Subscriber: &clientID, Topic: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Subscription status - : %t", r.GetDone())
}
