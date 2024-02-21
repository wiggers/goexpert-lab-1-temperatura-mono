package controller

import (
	"encoding/json"
	"net/http"

	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/entity"
	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/infra/adapter"
	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/usecase"
)

type findTemperatureByZipCodeController struct {
	City    entity.CityAdapterInterface
	Weather entity.WeatherAdapterInterface
}

func NewFindTemperatureByZipCodeController() *findTemperatureByZipCodeController {

	city := adapter.NewBrasilApiData()
	weather := adapter.NewWeatherApi()

	return &findTemperatureByZipCodeController{
		City:    city,
		Weather: weather,
	}
}

func (f *findTemperatureByZipCodeController) FindTemperature(w http.ResponseWriter, r *http.Request) {

	var dto usecase.ZipCodeInputDto
	dto.ZipCode = r.URL.Query().Get("zipcode")
	if dto.ZipCode == "" {
		http.Error(w, "invalid zip code", http.StatusBadRequest)
		return
	}

	findTemperatureByZip := usecase.NewTemperatureByZipCode(f.City, f.Weather)
	response, err := findTemperatureByZip.Execute(dto)

	if err != nil {
		if err.Error() == "invalid zip code" {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		if err.Error() == "can not find zip code" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
