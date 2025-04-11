package main

import (
	"context"
	"fmt"
	"github.com/DKW2/MuCache_Extended/internal/acmeair"
	"github.com/DKW2/MuCache_Extended/internal/social"
	"net/http"
	"runtime"

	"github.com/DKW2/MuCache_Extended/pkg/wrappers"
)

func heartbeat(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Heartbeat\n"))
	if err != nil {
		return
	}
}

func getFlightsByAirportsAndDepartureDate(ctx context.Context, req *acmeair.GetFlightsRequest) *acmeair.GetFlightsResponse {
	flights := acmeair.GetFlightsByAirportsAndDepartureDate(ctx, req.FromAirport, req.ToAirport, req.DeptDate)
	//fmt.Printf("Flights read: %+v\n", flights)
	resp := acmeair.GetFlightsResponse{Flights: flights}
	return &resp
}

func createSegment(ctx context.Context, req *acmeair.CreateSegmentRequest) *acmeair.CreateSegmentResponse {
	segmentName := acmeair.CreateSegment(ctx, req.OriginPort, req.DestPort, req.Miles)
	//fmt.Println("Segment stored: " + segment)
	resp := acmeair.CreateSegmentResponse{
		FlightName: segmentName
	}
	return &resp
}


// TODO: Create entrypoint for createFlight

func main() {
	fmt.Println(runtime.GOMAXPROCS(8))
	http.HandleFunc("/heartbeat", heartbeat)
	http.HandleFunc("/create_segment", wrappers.NonROWrapper[acmeair.CreateSegmentRequest, acmeair.CreateSegmentResponse](createSegment))
	http.HandleFunc("/ro_get_flights", wrappers.ROWrapper[acmeair.GetFlightsRequest, acmeair.GetFlightsResponse](getFlightsByAirportsAndDepartureDate))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
