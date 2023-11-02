package testpack

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	bus_pb "github.com/bharad1988/instantbus/busproto"
	pb "github.com/bharad1988/instantbus/subscriber"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestFunc() {
	fmt.Println("Called, it works , isn't it !")
}

var (
	addr = flag.String("addr", "localhost:5001", "the address to connect to")
	port int
)

// Push service (listen to topic)
type server struct {
	pb.UnimplementedPushServer
}

type SubscriptionService struct {
	Conn *grpc.ClientConn
	once sync.Once
}

// Handle pushed messages
func (s *server) PushMessage(ctx context.Context, in *pb.PushMessageRequest) (*pb.PushMessageReply, error) {
	log.Printf("𝙩𝙤𝙥𝙞𝙘 : %s :: 𝙢𝙚𝙨𝙨𝙖𝙜𝙚 :- %s", in.GetTopic(), in.GetMessage())
	e := time.Now().Unix()
	return &pb.PushMessageReply{Ts: &e}, nil
}

func (s *SubscriptionService) StartSubscriptionService(serviceBus string) {
	// defer conn.Close()
	start := func() {
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		s.Conn = conn
		rand.NewSource(time.Now().UnixNano())
		// generate random integer between 10 and 20
		port = rand.Intn(60000) + 1024
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
	go startListner(&s.once, start)
	time.Sleep(100 * time.Millisecond)
}

func startListner(once *sync.Once, task func()) {
	once.Do(task)
}

func (s *SubscriptionService) SubscribeTopic(topic string) {
	flag.Parse()
	// Set up a connection to the server.
	c := bus_pb.NewBusClient(s.Conn)
	clientID := fmt.Sprintf("localhost:%d", port)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r, err := c.SubTopic(ctx, &bus_pb.SubscribeRequest{Subscriber: &clientID, Topic: &topic})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Subscription status - : %t", r.GetDone())
}
