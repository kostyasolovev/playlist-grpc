package app

import (
	"context"
	"fmt"

	"github.com/kostyasolovev/playlist-grpc/config"
	"github.com/kostyasolovev/playlist-grpc/pkg/api"
	"github.com/kostyasolovev/playlist-grpc/src/ytplaylist"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"
)

type YoutubeGRPCServer struct {
	// youtube client
	ytService *youtube.Service
	// функция обработки респонсов от youtube api
	getFunc func(string) ([]string, error)
	// реализация интерфейса YoutubePlaylistServer.
	api.UnimplementedYoutubePlaylistServer
}

// реализация интерфейса YoutubePlaylistServer.
func (grpcServ *YoutubeGRPCServer) List(ctx context.Context, r *api.PlaylistRequest) (resp *api.PlaylistResponse, err error) {
	// GET request to youtube api
	ans, err := grpcServ.getFunc(r.Id)
	if err != nil {
		if gerr, ok := err.(*googleapi.Error); !ok || gerr.Code != 404 {
			resp.Err = "internal server error"
		} else {
			// 404 statusNotFound
			resp.Err = fmt.Sprintf("playlist with Id [%s] not found", r.Id)
		}
	}

	resp.Item = append(resp.Item, ans...)

	return resp, err
}

// настраиваем YoutubeGRPCServer.
func (grpcServ *YoutubeGRPCServer) Setup(ctx context.Context, cfg *config.Config) error {
	// youtube api client
	service, err := ytplaylist.NewYTServiceWithApiKey(ctx, cfg.YoutubeApiKey)
	if err != nil {
		return err
	}

	grpcServ.ytService = service
	// youtube response handler
	grpcServ.getFunc = func(playlistId string) ([]string, error) {
		var ansLimit int64 = 10 // лимит элементов

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
			// нас интересуют только айдишники
			ans = append(ans, v.ContentDetails.VideoId)
		}

		return ans, nil
	}

	return nil
}
