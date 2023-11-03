package sbclient

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

type queueData struct {
	messages [][]byte
	index int
}

// Push service (listen to topic)
type server struct {
	pb.UnimplementedPushServer
	TopicQueue map[string]queueData
}

type SubscriptionService struct {
	Conn   *grpc.ClientConn
	once   sync.Once
	Client bus_pb.BusClient
	subServer server
}

// Handle pushed messages
func (s *server) PushMessage(ctx context.Context, in *pb.PushMessageRequest) (*pb.PushMessageReply, error) {
	// log.Printf("𝙩𝙤𝙥𝙞𝙘 : %s :: 𝙢𝙚𝙨𝙨𝙖𝙜𝙚 :- %s", in.GetTopic(), in.GetMessage())
	messageData := s.TopicQueue[in.GetTopic()]
	messageData.messages = append(messageData.messages, in.GetMessage())
	s.TopicQueue[in.GetTopic()] = messageData
	e := time.Now().Unix()
	// log.Printf("topic - %s , message %s ", in.GetTopic(),s.TopicQueue[in.GetTopic()])

	return &pb.PushMessageReply{Ts: &e}, nil
}

func (s *SubscriptionService) StartPublisher(serviceBus string) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	s.Conn = conn
	s.Client = bus_pb.NewBusClient(s.Conn)

}

func (s *SubscriptionService) StartSubscriptionService(serviceBus string) {
	// defer conn.Close()
	start := func() {
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		s.Conn = conn
		s.Client = bus_pb.NewBusClient(s.Conn)
		rand.NewSource(time.Now().UnixNano())
		// generate random integer between 10 and 20
		port = rand.Intn(60000) + 1024
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		s.subServer = server{}
		s.subServer.TopicQueue = make(map[string]queueData)
		pb.RegisterPushServer(grpcServer, &s.subServer)
		log.Printf("server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
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
	clientID := fmt.Sprintf("localhost:%d", port)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	r, err := s.Client.SubTopic(ctx, &bus_pb.SubscribeRequest{Subscriber: &clientID, Topic: &topic})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Subscription status - : %t", r.GetDone())
}

func (s *SubscriptionService) GetAllMessages(topic string) [][]byte{
	return s.subServer.TopicQueue[topic].messages
}

func (s *SubscriptionService) GetAllUnreadMessages(topic string) [][]byte {
	queueData :=  s.subServer.TopicQueue[topic]
	l := len(queueData.messages)
	slice := queueData.messages[s.subServer.TopicQueue[topic].index:l]
	queueData.index = l
	s.subServer.TopicQueue[topic] = queueData
	return slice
}

func (s *SubscriptionService) SendMessage(topic string, message []byte) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := s.Client.SendMessage(ctx, &bus_pb.MessageRequest{Topic: &topic, Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Sent at: %d", r.GetTs())

}
