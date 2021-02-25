package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math"
	"net"
	pb "end/proto"
)

type server struct{}

func (s *server) PrimeNumberDecomposition(ctx context.Context, in *pb.PrimeRequest) (*pb.PrimeReply, error) {
	n := in.N1
	var ans []int
	for ;n%2==0;n=n/2 {
		fmt.Print(2, " ")
		ans = append(ans, 2)
	}
	i := 3
	for ;i<=int(math.Sqrt(float64(n))); i = i + 2 {
		for ;n%i==0; {
			fmt.Print(i, " ")
			ans = append(ans, i)
			n = n/i
		}
	}

	if n > 2 {
		fmt.Print(n, " ")
		ans = append(ans, n)
	}
	&pb.PrimeReply{N1: ans}, nil
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50050")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Server is running on port:50050")
	if err := s.Server(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
