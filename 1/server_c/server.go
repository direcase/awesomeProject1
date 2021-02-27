package main

import (
	"awesomeProject1/dariyaproto"

	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Server struct {
	dariyaproto.UnimplementedPrimeServiceServer
}

func (s *Server) Do(req *dariyaproto.Request, stream dariyaproto.PrimeService_DoServer) error {
	number := req.GetNumber()
	var a []int = Prime(int(number))

	for i := 0; i < len(a); i++ {

		res := &dariyaproto.Response{Prime: int32(a[i])}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending responses: %v", err.Error())
		}
		time.Sleep(time.Second)
	}

	return nil
}
func Prime(n int) (rest []int) {
	for n%2 == 0 {
		rest = append(rest, 2)
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			rest = append(rest, i)
			n = n / i
		}
	}
	if n > 2 {
		rest = append(rest, n)
	}

	return
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	dariyaproto.RegisterPrimeServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
