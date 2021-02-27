package main

import (
	"awesomeProject1/dariyaproto"
	_ "awesomeProject1/dariyaproto"
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
	fmt.Println(request)

}
