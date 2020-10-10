package server

import (
	"Go-grpc/pkg/errcode"
	"Go-grpc/pkg/tagApi"
	pb "Go-grpc/proto"
	"context"
	"encoding/json"
)

/**
* @Author: super
* @Date: 2020-09-25 08:33
* @Description:
**/

type TagServer struct {
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := tagApi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}
	return &tagList, nil
}