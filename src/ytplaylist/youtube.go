package ytplaylist

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Создает http клиента для youtube API.
func NewYTServiceWithApiKey(ctx context.Context, key string) (*youtube.Service, error) {
	service, err := youtube.NewService(ctx, option.WithAPIKey(key))
	if err != nil {
		return nil, errors.Wrap(err, "NewYTServiceWithApiKey")
	}

	return service, nil
}

// делает запрос на получение плейлиста к Youtube API.
func GetYoutubePlaylist(ctx context.Context, service *youtube.Service, id string, limit int64, parts ...string) (*youtube.PlaylistItemListResponse, error) {
	call := service.PlaylistItems.List(parts)
	call = call.MaxResults(limit)
	call = call.Context(ctx)
	call = call.PlaylistId(id)
	response, err := call.Do()
	if err != nil {
		return nil, errors.Wrap(err, "GetYoutubePlaylist")
	}

	return response, nil
}
