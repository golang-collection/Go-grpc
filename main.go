package main

import (
	"Go-grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "Go-grpc/proto"
)

/**
* @Author: super
* @Date: 2020-09-22 19:30
* @Description:
**/

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	listener, err := net.Listen("tcp", ":5000")
	if err != nil{
		log.Fatalf("net.Listen err: %v\n", err)
		return
	}

	err = s.Serve(listener)
	if err != nil{
		log.Fatalf("server.Server err: %v\n", err)
	}

}