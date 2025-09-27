package weather

type Client struct {
	apiURL string
}

type Service struct {
	client *Client
}

type WeatherResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Current   struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
	} `json:"current_weather"`
}