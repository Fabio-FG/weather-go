package weather

func NewService (client *Client) *Service {
	return &Service{client: client}
}

func (s *Service) GetWeather(lat, lon float64) (*WeatherResponse, error){
	return s.client.GetWeather(lat, lon)
}