package app

import (
	"context"
	"grpc-server/config"
	"grpc-server/pkg/api"
	"grpc-server/src/ytplaylist"

	"google.golang.org/api/youtube/v3"
)

type YoutubeGRPCServer struct {
	ytService *youtube.Service
	getFunc   func(string) ([]string, error)
}

func (grpcServ *YoutubeGRPCServer) List(ctx context.Context, r *api.PlaylistRequest) (*api.PlaylistResponse, error) {
	// items, err := grpcServ.getFunc(fmt.Sprintf(""))
	return &api.PlaylistResponse{}, nil
}

func (grpcServ *YoutubeGRPCServer) Setup(ctx context.Context, cfg *config.Config) error {
	service, err := ytplaylist.NewYTServiceWithApiKey(ctx, cfg.YoutubeApiKey)
	if err != nil {
		return err
	}

	grpcServ.ytService = service
	grpcServ.getFunc = func(playlistId string) ([]string, error) {
		var ansLimit int64 = 10

		ytResponse, err := ytplaylist.GetYoutubePlaylist(ctx, grpcServ.ytService,
			playlistId, ansLimit, "snippet", "contentDetails")
		if err != nil {
			return nil, err
		}

		ans := make([]string, 0, ansLimit)

		for i, v := range ytResponse.Items {
			if i >= int(ansLimit)-1 {
				break
			}

			ans = append(ans, v.ContentDetails.VideoId)
		}

		return ans, nil
	}

	return nil
}
