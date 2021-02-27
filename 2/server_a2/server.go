package main

import (
	"awesomeProject1/dproto2"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type Server struct {
	dproto2.UnimplementedPrimeServiceServer
}

func (s *Server) Do(stream dproto2.PrimeService_DoServer) error {
	var result float64 = 0
	var number []int32
	sum := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&dproto2.Response{Avg: result,
			})

		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		number = req.GetNumber()
		for i := 0; i < len(number); i++ {
			sum += int(number[i])
		}
		result = float64(sum / len(number))
		//result = result+ (number)
	}
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	dproto2.RegisterPrimeServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
