package main

import (
	"context"
	"fmt"
	"github.com/DKW2/MuCache_Extended/internal/social"
	"github.com/DKW2/MuCache_Extended/pkg/cm"
	"github.com/DKW2/MuCache_Extended/pkg/common"
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

func ComposePost(ctx context.Context, req *social.ComposePostRequest) *string {
	social.ComposePost(ctx, req.Text, req.CreatorId)
	resp := "OK"
	return &resp
}

func ComposePostMulti(ctx context.Context, req *social.ComposePostMultiRequest) *string {
	social.ComposeMulti(ctx, req.Text, req.Number, req.CreatorId)
	resp := "OK"
	return &resp
}

func main() {
	if common.ShardEnabled {
		fmt.Println(runtime.GOMAXPROCS(1))
	} else {
		fmt.Println(runtime.GOMAXPROCS(8))
	}
	go cm.ZmqProxy()
	http.HandleFunc("/heartbeat", heartbeat)
	http.HandleFunc("/compose_post", wrappers.NonROWrapper[social.ComposePostRequest, string](ComposePost))
	http.HandleFunc("/compose_post_multi", wrappers.NonROWrapper[social.ComposePostMultiRequest, string](ComposePostMulti))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
