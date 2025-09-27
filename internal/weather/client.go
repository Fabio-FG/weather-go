package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewClient(apiURL string) *Client {
	return &Client{apiURL: apiURL}
}

func (c *Client) GetWeather(lat, lon float64) (*WeatherResponse, error){
	url := fmt.Sprintf("%s?latitude=%f&longitude=%f&current_weather=true", c.apiURL, lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var data WeatherResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}