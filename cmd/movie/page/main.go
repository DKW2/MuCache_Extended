package main

import (
	"context"
	"fmt"
	"github.com/DKW2/MuCache_Extended/internal/movie"
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

func readPage(ctx context.Context, req *movie.ReadPageRequest) *movie.ReadPageResponse {
	page := movie.ReadPage(ctx, req.MovieId)
	//fmt.Printf("Page read: %v\n", page)
	resp := movie.ReadPageResponse{Page: page}
	return &resp
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(8))
	go cm.ZmqProxy()
	http.HandleFunc("/heartbeat", heartbeat)
	http.HandleFunc("/ro_read_page", wrappers.ROWrapper[movie.ReadPageRequest, movie.ReadPageResponse](readPage))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
