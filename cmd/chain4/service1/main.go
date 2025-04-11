package main

import (
	"context"
	"fmt"
	"github.com/DKW2/MuCache_Extended/internal/loadcm"
	"github.com/DKW2/MuCache_Extended/internal/twoservices"
	"github.com/DKW2/MuCache_Extended/pkg/cm"
	"github.com/DKW2/MuCache_Extended/pkg/invoke"
	"github.com/DKW2/MuCache_Extended/pkg/wrappers"
	"math/rand"
	"net/http"
	"runtime"
)

var Callee = "service2"
var MaxProcs = 8

func heartbeat(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Heartbeat\n"))
	if err != nil {
		return
	}
}

func read(ctx context.Context, req *twoserivces.ReadRequest) *twoserivces.ReadResponse {
	resp := invoke.Invoke[twoserivces.ReadResponse](ctx, Callee, "ro_read", req)
	return &resp
}

func write(ctx context.Context, req *twoserivces.WriteRequest) *string {
	resp := invoke.Invoke[string](ctx, Callee, "write", req)
	return &resp
}

func hitormiss(ctx context.Context, req *twoserivces.HitOrMissRequest) *string {
	dice := rand.Float32()
	if dice < req.HitRate {
		invoke.InvokeHit(ctx, Callee, "ro_hitormiss", req)
	} else {
		invoke.InvokeMiss[string](ctx, Callee, "ro_hitormiss", req)
	}
	resp := "OK"
	return &resp
}

func invalidationExperiment(ctx context.Context, req *loadcm.InvalidationExperimentRequest) *string {
	// Start running the zmqfeeder
	fmt.Printf("Starting experiment for: %v \n", req.Times)
	go twoserivces.InvalidationExperiment(req.Times, req.Timeout, Callee, "ro_read", "backend", "write")
	resp := "OK"
	return &resp
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(MaxProcs))
	go cm.ZmqProxy()
	http.HandleFunc("/heartbeat", heartbeat)
	http.HandleFunc("/ro_read", wrappers.ROWrapper[twoserivces.ReadRequest, twoserivces.ReadResponse](read))
	http.HandleFunc("/write", wrappers.NonROWrapper[twoserivces.WriteRequest, string](write))
	http.HandleFunc("/ro_hitormiss", wrappers.ROWrapper[twoserivces.HitOrMissRequest, string](hitormiss))
	http.HandleFunc("/invalidation_experiment", wrappers.NonROWrapper[loadcm.InvalidationExperimentRequest, string](invalidationExperiment))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
