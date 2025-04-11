package main

import (
	"context"
	"fmt"
	"github.com/DKW2/MuCache_Extended/internal/boutique"
	"github.com/DKW2/MuCache_Extended/pkg/cm"
	"github.com/DKW2/MuCache_Extended/pkg/wrappers"
	"net/http"
	"runtime"
)

func heartbeat(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Heartbeat\n"))
	if err != nil {
		return
	}
}

func getRecommendations(ctx context.Context, req *boutique.GetRecommendationsRequest) *boutique.GetRecommendationsResponse {
	products := boutique.GetRecommendations(ctx, req.ProductIds)
	resp := boutique.GetRecommendationsResponse{ProductIds: products}
	return &resp
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(8))
	go cm.ZmqProxy()
	http.HandleFunc("/heartbeat", heartbeat)
	http.HandleFunc("/ro_get_recommendations", wrappers.ROWrapper[boutique.GetRecommendationsRequest, boutique.GetRecommendationsResponse](getRecommendations))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
