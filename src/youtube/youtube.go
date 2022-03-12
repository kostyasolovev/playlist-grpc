package youtube

import "context"

type Item struct {
	Name string
}

func GetPlaylistWithTimeout(ctx context.Context, url string) ([]Item, error) {
	ans := []Item{}

	return ans, nil
}
