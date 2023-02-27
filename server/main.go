package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc-load-balancing/pb"
	"net"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
	addr string
}

func (g greeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	//TODO implement me
	return &pb.HelloReply{Message: g.addr}, nil
}

func main() {
	port := flag.String("port", "8080", "")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", *port))
	if err != nil {
		panic(err)
	}
	fmt.Printf("grpc run port: %s\n", *port)
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &greeterServer{addr: *port})
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
