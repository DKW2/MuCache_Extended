package movie

import (
	"context"
	"github.com/DKW2/MuCache_Extended/pkg/state"
)

func RegisterMovieId(ctx context.Context, title string, movieId string) {
	state.SetState(ctx, title, movieId)
}

func GetMovieId(ctx context.Context, title string) string {
	movieId, err := state.GetState[string](ctx, title)
	if err != nil {
		panic(err)
	}
	return movieId
}
