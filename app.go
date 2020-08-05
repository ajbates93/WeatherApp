package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

// func GetIlkleyWeather() {

// }

func main() {

	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=ilkley,uk&appid=11a2659c258f169b1711290aafc9c04d")

	if err != nil {
		log.Error(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Info(string(responseData))

	log.Info("Starting WeatherApp API Server")
	router := mux.NewRouter()
	http.ListenAndServe(":8000", router)
}
