package main

import (
	"context"
	"fmt"
	"github.com/DKW2/MuCache_Extended/internal/hotel"
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

func storeRate(ctx context.Context, req *hotel.StoreRateRequest) *hotel.StoreRateResponse {
	hotelId := hotel.StoreRate(ctx, req.Rate)
	//fmt.Println("Movie info stored for id: " + movieId)
	resp := hotel.StoreRateResponse{HotelId: hotelId}
	return &resp
}

func getRates(ctx context.Context, req *hotel.GetRatesRequest) *hotel.GetRatesResponse {
	rates := hotel.GetRates(ctx, req.HotelIds)
	//fmt.Printf("Movie info read: %v\n", movieInfo)
	resp := hotel.GetRatesResponse{Rates: rates}
	//fmt.Printf("[ReviewStorage] Response: %v\n", resp)
	return &resp
}

func main() {
	fmt.Println(runtime.GOMAXPROCS(8))
	go cm.ZmqProxy()
	http.HandleFunc("/heartbeat", heartbeat)
	http.HandleFunc("/store_rate", wrappers.NonROWrapper[hotel.StoreRateRequest, hotel.StoreRateResponse](storeRate))
	http.HandleFunc("/ro_get_rates", wrappers.ROWrapper[hotel.GetRatesRequest, hotel.GetRatesResponse](getRates))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
