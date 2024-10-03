package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var quotes = []string{
	"People Are Gonna Say The Targets On Our Back But I Hope It’s Right On Our Forehead Between Our Eyes",
	"There's no such thing as a foul, you either die or you don't",
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomResponse := quotes[rng.Intn(len(quotes))]

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&randomResponse); err != nil {
			log.Printf("Error encoding JSON response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	})

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}
