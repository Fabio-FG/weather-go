package main

import (
	"log"
	"net/http"
	"os"
	"weather-go/internal/weather"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	apiURL := os.Getenv("API_URL")
	client := weather.NewClient(apiURL)
	service := weather.NewService(client)

	http.HandleFunc("/weather", weather.WeatherHandler(service))

	log.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
