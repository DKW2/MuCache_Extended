package main

import (
	"github.com/DKW2/MuCache_Extended/pkg/cm"
	"net/http"
)

func heartbeat(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Heartbeat\n"))
	if err != nil {
		return
	}
}

func main() {
	go cm.ZmqProxy()
	http.HandleFunc("/heartbeat", heartbeat)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
