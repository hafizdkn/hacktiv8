package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"project_3/helper"
)

type Service interface {
	GetLocation() *helper.Response
}
type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) GetLocation() *helper.Response {
	var dataLocation Location
	url := "https://geolocation.onetrust.com/cookieconsentpub/v1/geo/location"
	resp, err := http.Get(url)
	if err != nil {
		return helper.ErrorResponse(err)
	}
	defer resp.Body.Close()
	rawData, _ := ioutil.ReadAll(resp.Body)
	data := cleanString(rawData)
	err = json.Unmarshal([]byte(*data), &dataLocation)
	if err != nil {
		return helper.ErrorResponse(err)
	}
	if dataLocation.City == "" {
		dataLocation.City = "surabaya"
	}

	responseWeather := generateDataWeahter(dataLocation)
	return responseWeather
}

func generateDataWeahter(location Location) *helper.Response {
	dataWeather, err := getWeather(location.City)
	if err != nil {
		return helper.ErrorResponse(err)
	}

	conditon := randomValueCondition()
	statusCondition := createStatusCondition(conditon)
	nextRefreshWeather := 10 * time.Minute

	weahterResponse := WeatherResponse{
		Location: location, Weather: dataWeather, Condition: conditon, Status: statusCondition, Time: nextRefreshWeather,
	}
	response := FormatWeatherResponse(weahterResponse)

	// save record weather to json
	writeDataToJson(response)

	res := helper.SuccessResponse(response, "SUCCESS GET DATA LOCATION")
	return res
}

func getWeather(city string) (Weather, error) {
	var weather Weather
	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s&appid=ce7a971952c7df04c4224435da5d818f&units=metric",
		city,
	)
	resp, err := http.Get(url)
	if err != nil {
		return weather, err
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &weather)
	if err != nil {
		return weather, err
	}
	weather.Main.Temp = int(weather.Main.Temp.(float64))
	return weather, nil
}

func randomValueCondition() Condition {
	v := make([]int, 2)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 2; i++ {
		v[i] = rand.Intn(100)
	}
	water := v[0]
	wind := v[1]
	return Condition{Water: water, Wind: wind}
}

func createStatusCondition(condition Condition) Status {
	statusWater := "aman"
	statusWind := "aman"

	if condition.Water > 8 {
		statusWater = "Bahaya"
	} else if condition.Water > 6 {
		statusWater = "Siaga"
	}

	if condition.Wind > 15 {
		statusWind = "Bahaya"
	} else if condition.Wind > 7 {
		statusWind = "Siaga"
	}
	return Status{StatusWater: statusWater, StatusWind: statusWind}
}

func writeDataToJson(data interface{}) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("weather.json", file, 0o644)
}

func cleanString(rawData []byte) *string {
	data := string(rawData)
	startSub := strings.Index(data, "(")
	stopSub := len(data) - 2
	newData := data[startSub+1 : stopSub]
	return &newData
}
