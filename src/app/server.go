package app

import (
	"context"
	"fmt"

	api "github.com/kostyasolovev/youtube_pb_api"
	"github.com/pkg/errors"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/youtube/v3"

	"playlist-grpc/config"
	"playlist-grpc/src/ytplaylist"
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
func (grpcServ *YoutubeGRPCServer) List(ctx context.Context, r *api.PlaylistRequest) (*api.PlaylistResponse, error) {
	resp := new(api.PlaylistResponse)
	// GET request to youtube api
	ans, err := grpcServ.getFunc(r.Id)
	if err != nil { // nolint: errorlint
		if gerr, ok := err.(*googleapi.Error); !ok || gerr.Code != 404 {
			resp.Err = "internal server error"
		} else {
			// 404 statusNotFound
			resp.Err = fmt.Sprintf("playlist with Id [%s] not found", r.Id)
			err = nil
		}
	}

	resp.Item = append(resp.Item, ans...)

	return resp, err
}

// настраиваем YoutubeGRPCServer.
func (grpcServ *YoutubeGRPCServer) Setup(ctx context.Context, cfg *config.Config) error {
	// youtube api client
	service, err := ytplaylist.NewYTServiceWithAPIKey(ctx, cfg.YoutubeAPIKey)
	if err != nil {
		return errors.Wrap(err, "creating youtube service failed")
	}

	grpcServ.ytService = service
	// youtube response handler
	grpcServ.getFunc = func(playlistId string) ([]string, error) {
		var ansLimit int64 = 10 // лимит элементов

		ytResponse, err := ytplaylist.GetYoutubePlaylist(ctx, grpcServ.ytService,
			playlistId, ansLimit, "snippet", "contentDetails")
		if err != nil {
			return nil, err // nolint: wrapcheck
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
