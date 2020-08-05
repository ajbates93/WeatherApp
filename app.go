package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func GetIlkleyWeather() {

}

func main() {
	log.Info("Starting WeatherApp API Server")
	router := mux.NewRouter()
	http.ListenAndServe(":8000", router)
}
