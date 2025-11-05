package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"weather-go/internal/weather"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var cityCmd = &cobra.Command{
	Use: "city [name]",
	Short: "Shows the realtime weather of a city",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cityName := strings.ToLower((args[0]))

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error while loading .env")
		}

		apiURL := os.Getenv("API_URL")
		if apiURL == "" {
			return fmt.Errorf("API_URL not found.")
		}

		client := weather.NewClient(apiURL)
		service := weather.NewService(client)

		coords, ok := weather.Cities[cityName]

		if !ok {
			return fmt.Errorf("city '%s' not found", cityName)
		}

		data, err := service.GetWeather(coords.Latitude, coords.Longitude)

		if err != nil {
			return fmt.Errorf("error while fetching data: %v", err)
		}

		fmt.Printf("🌤️ Weather in %s (lat: %.2f, lon: %.2f)\n", coords.Name, coords.Latitude, coords.Longitude)
		fmt.Printf("Temperature: %.1f°C\n", data.Current.Temperature)
		fmt.Printf("Wind: %.1f km/h\n", data.Current.WindSpeed)
		return nil
	},
}

func init(){
	rootCmd.AddCommand(cityCmd)
}