package main

import (
	"context"
	"google.golang.org/grpc"
	"log"

	pb "Go-grpc/proto"
)

/**
* @Author: super
* @Date: 2020-10-10 15:36
* @Description:
**/

func main() {
	ctx := context.Background()
	conn, err := GetClientConn(ctx, "localhost:5000", nil)
	if err != nil{
		log.Fatalf("err: %v", err)
	}
	defer conn.Close()

	tagServiceClient := pb.NewTagServiceClient(conn)
	resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})
	if err != nil{
		log.Fatalf("tagServiceClient.GetTagList err: %v", err)
	}
	log.Printf("resp:%v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}