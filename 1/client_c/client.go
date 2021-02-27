package main

import (
	"awesomeProject1/dariyaproto"
	_ "awesomeProject1/dariyaproto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello I'm a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := dariyaproto.NewPrimeServiceClient(conn)
	PrimeNumberDecomposition(c)
}

func PrimeNumberDecomposition(c dariyaproto.PrimeServiceClient) {

	request := dariyaproto.Request{Number: 120}
	ctx := context.Background()
	response, err := c.Do(ctx, request)
	if err != nil {
		log.Fatalf("error while calling Greet RPC %v", err)
	}
	log.Printf("response from Greet:%v", response)

}
