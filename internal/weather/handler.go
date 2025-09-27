package weather

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func WeatherHandler(service *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		latStr := r.URL.Query().Get("lat")
		lonStr := r.URL.Query().Get("lon")

	
		lat, err1 := strconv.ParseFloat(latStr, 64)
		lon, err2 := strconv.ParseFloat(lonStr, 64)

		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid latitude or longitude", http.StatusBadRequest)
			return;
		}

		data, err := service.GetWeather(lat, lon)

		if err != nil {
			http.Error(w, "Error fetchign weather", http.StatusInternalServerError)
			return;
		}

		json.NewEncoder(w).Encode(data)
	}
}

