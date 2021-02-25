package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
	pb "end/proto"
	"google.golang.org/grpc"
)

func argParser(n1 string) (int32) {
	N1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Cannot parse arge[1]: %s", err)
	}
	return int32(N1)
}

func main() {
	if len(os.Args) != 2 {
            log.Fatalf("1 number expected: n1")
	}

	n1 := argParser(os.Args[1])

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	Result, err := client.Add(ctx, &pb.PrimeRequest{int32(n1)})
	if err != nil {
		log.Fatalf("Adding error: %s", err)
	}
	log.Printf("Result: %d", Result.N1)
}
