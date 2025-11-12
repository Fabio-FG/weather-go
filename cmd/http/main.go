package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"weather-go/internal/weather"
)

func main() {
	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		log.Fatal("API_URL environment variable not set")
	}

	client := weather.NewClient(apiURL)
	service := weather.NewService(client)

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "Missing ?city= parameter", http.StatusBadRequest)
			return
		}

		coords, ok := weather.Cities[city]
		if !ok {
			http.Error(w, "City not found", http.StatusNotFound)
			return
		}

		data, err := service.GetWeather(coords.Latitude, coords.Longitude)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "🌤️ Weather in %s: %.1f°C (Wind: %.1f km/h)",
			coords.Name, data.Current.Temperature, data.Current.WindSpeed)
	})

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
