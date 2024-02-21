package main

import (
	"net/http"

	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/infra/controller"
)

func main() {
	//config := configs.LoadConfig(".")

	controller := controller.NewFindTemperatureByZipCodeController()

	http.HandleFunc("/temperature", controller.FindTemperature)
	http.ListenAndServe(":8080", nil)

}
