package weather

import "time"

type Condition struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type Status struct {
	StatusWater string `json:"status_water"`
	StatusWind  string `json:"status_wind"`
}

type Location struct {
	City      string `json:"city"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type WeatherResponse struct {
	Location  Location      `json:"location"`
	Weather   Weather       `json:"weather"`
	Condition Condition     `json:"condition"`
	Status    Status        `json:"status"`
	Time      time.Duration `json:"time"`
}

type Weather struct {
	Base   string `json:"base"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Cod   int `json:"cod"`
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Dt   int `json:"dt"`
	ID   int `json:"id"`
	Main struct {
		FeelsLike float64     `json:"feels_like"`
		Humidity  int         `json:"humidity"`
		Pressure  int         `json:"pressure"`
		Temp      interface{} `json:"temp"`
		TempMax   float64     `json:"temp_max"`
		TempMin   float64     `json:"temp_min"`
	} `json:"main"`
	Name string `json:"name"`
	Rain struct {
		Hour float64 `json:"1h"`
	} `json:"rain"`
	Sys struct {
		Country string `json:"country"`
		ID      int    `json:"id"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
		Type    int    `json:"type"`
	} `json:"sys"`
	Timezone   int `json:"timezone"`
	Visibility int `json:"visibility"`
	Weather    []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
		ID          int    `json:"id"`
		Main        string `json:"main"`
	} `json:"weather"`
	Wind struct {
		Deg   int     `json:"deg"`
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func FormatWeatherResponse(weahter WeatherResponse) WeatherResponse {
	return WeatherResponse{
		Location: weahter.Location, Weather: weahter.Weather, Condition: weahter.Condition, Status: weahter.Status,
	}
}
