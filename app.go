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

func GetIlkleyWeather(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=ilkley,uk&appid=11a2659c258f169b1711290aafc9c04d")

	if err != nil {
		log.Error("The HTTP Request failed with error \n", err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("The HTTP Request failed with error \n", err)
	}

	log.Info(string(responseData))
}

func main() {

	log.Info("Starting WeatherApp API Server")
	router := mux.NewRouter()
	router.HandleFunc("/ilkley", GetIlkleyWeather).Methods("GET")
	http.ListenAndServe(":8000", router)
}
