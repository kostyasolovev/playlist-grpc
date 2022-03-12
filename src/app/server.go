package app

import (
	"context"
	"grpc-server/pkg/api"
	"grpc-server/src/youtube"
)

type YoutubeGRPCServer struct {
	getFunc func(context.Context, string) ([]youtube.Item, error)
}

func (grpcServ *YoutubeGRPCServer) List(ctx context.Context, r *api.PlaylistRequest) (*api.PlaylistResponse, error) {
	// items, err := grpcServ.getFunc(fmt.Sprintf(""))
	return &api.PlaylistResponse{}, nil
}
