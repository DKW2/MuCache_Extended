package movie

import (
	"context"
	"github.com/DKW2/MuCache_Extended/pkg/state"
)

func WritePlot(ctx context.Context, plotId string, plot string) string {
	state.SetState(ctx, plotId, plot)
	return plotId
}

func ReadPlot(ctx context.Context, plotId string) string {
	plot, err := state.GetState[string](ctx, plotId)
	if err != nil {
		panic(err)
	}
	return plot
}
