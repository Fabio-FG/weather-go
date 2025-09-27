package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"weather-go/internal/weather"

	"github.com/joho/godotenv"
)

func main() {
	city := flag.String("city", "Lisbon", "City name (e.g. Lisbon, Tokyo, New York)")

	lat := flag.Float64("lat", 38.7169, "Latitude (Lisbon default)")
	lon := flag.Float64("lon", -9.139, "Longitude (Lisbon default)")

	flag.Parse()

	if *city == "" {
		fmt.Println("No city provided, defaulting to Lisbon coords")
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	apiURL := os.Getenv("API_URL")
	client := weather.NewClient(apiURL)
	service := weather.NewService(client)

	cityKey := strings.ToLower(*city)
	
	coords, ok := weather.Cities[cityKey]
	data, err := service.GetWeather(coords.Latitude, coords.Longitude)

	if !ok {
		log.Fatalf("City '%s' not found in local map", *city)
	}

	if err != nil {
		log.Fatal("Error fetching weather:", err)
	}

	fmt.Printf("🌤️  Weather for %s (lat: %.2f, lon: %.2f)\n", *city, *lat, *lon)
	fmt.Printf("Temperature: %.1f°C\n", data.Current.Temperature)
	fmt.Printf("Wind Speed: %.1f km/h\n", data.Current.WindSpeed)
}
