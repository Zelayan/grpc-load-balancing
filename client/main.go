package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-load-balancing/pb"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	hello, err := c.SayHello(ctx, &pb.HelloRequest{Name: "cc"})
	if err != nil {
		panic(err)
	}
	fmt.Print(hello.GetMessage())
}
