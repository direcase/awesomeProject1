package main

import (
	"awesomeProject1/dproto2"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := dproto2.NewPrimeServiceClient(conn)
	GetNumbers(c)
}

func GetNumbers(c dproto2.PrimeServiceClient) {
	requests := []*dproto2.Request{
		{
			Number: 2,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
	}

	ctx := context.Background()
	stream, err := c.Do(ctx)
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf(" Response: %v\n", res)
}
