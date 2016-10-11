package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Locations is locations
type Locations []*Location

// Location is a location
type Location struct {
	Name string
	ID   string
	mod  float32
}

// Prices are prices
type Prices []Price

// Price is a price
type Price struct {
	Location  *Location
	Timestamp time.Time
	Price     float32
}

func locations() Locations {
	return []*Location{
		{Name: "Unit A", ID: "UA", mod: 1.0},
		{Name: "Unit B", ID: "UB", mod: 2.0},
		{Name: "Unit Z.1", ID: "UZ1", mod: 3.0},
		{Name: "Unit Z.2", ID: "UZ2", mod: 4.0},
	}
}

func createPrices(locations Locations) Prices {

	now := time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC)
	prices := Prices{}

	for i, location := range locations {
		for t := 0; t < 24; t = t + 1 {
			offset := time.Duration(t) * time.Hour
			price := Price{
				Location:  location,
				Timestamp: now.Add(offset),
				Price:     location.mod * float32(i*t),
			}

			prices = append(prices, price)
		}
	}

	return prices
}

func (locations Locations) endPoint(w http.ResponseWriter, r *http.Request) {

	if failure(w, r) == true {
		return
	}

	body, _ := json.Marshal(locations)
	fmt.Fprintf(w, string(body))
}

func (prices Prices) endPoint(w http.ResponseWriter, r *http.Request) {
	counter = counter + 1

	if failure(w, r) == true {
		return
	}

	id := r.URL.Query().Get("location_id")
	found := Prices{}

	for _, price := range prices {
		if price.Location.ID == id {
			found = append(found, price)
		}
	}

	body, _ := json.Marshal(found)
	fmt.Fprintf(w, string(body))
}

func failure(w http.ResponseWriter, r *http.Request) bool {
	n := rand.Intn(8)
	if n > 6 {
		time.Sleep(5 * time.Second)
	}

	counter = counter + 1
	if counter == 3 {
		counter = 0
		http.Error(w, "Unhelpful error message", http.StatusInternalServerError)
		return true
	}

	return false
}

var counter = 0

func main() {
	locs := locations()
	prices := createPrices(locs)

	http.HandleFunc("/locations", locs.endPoint)
	http.HandleFunc("/prices", prices.endPoint)

	fmt.Print("Starting the Pharos Code Eval server.  Happy Coding.\n")
	http.ListenAndServe(":4000", nil)
}
