package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/bharad1988/instantbus/busproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "earth"
)

var (
	addr = flag.String("addr", "localhost:5001", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBusClient(conn)
	for i := 0; i < 100; i++ {
		m := fmt.Sprintf("test %d", i)
		testMsg := []byte(m)
		// Contact the server and print out its response.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SendMessage(ctx, &pb.MessageRequest{Topic: name, Message: testMsg})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Sent at: %d", r.GetTs())
	}
}
