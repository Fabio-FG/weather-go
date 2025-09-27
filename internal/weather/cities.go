package weather

type City struct {
    Name      string
    Latitude  float64
    Longitude float64
}

var Cities = map[string]City{
    "lisbon":   {Name: "Lisbon", Latitude: 38.7169, Longitude: -9.139},
    "newyork":  {Name: "New York", Latitude: 40.7128, Longitude: -74.0060},
    "london":   {Name: "London", Latitude: 51.5074, Longitude: -0.1278},
    "tokyo":    {Name: "Tokyo", Latitude: 35.6895, Longitude: 139.6917},
    "paris":    {Name: "Paris", Latitude: 48.8566, Longitude: 2.3522},
    "sydney":   {Name: "Sydney", Latitude: -33.8688, Longitude: 151.2093},
    "berlin":   {Name: "Berlin", Latitude: 52.5200, Longitude: 13.4050},
    "madrid":   {Name: "Madrid", Latitude: 40.4168, Longitude: -3.7038},
    "dubai":    {Name: "Dubai", Latitude: 25.276987, Longitude: 55.296249},
    "singapore": {Name: "Singapore", Latitude: 1.3521, Longitude: 103.8198},
}