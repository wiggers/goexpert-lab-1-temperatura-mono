package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/wiggers/goexpert/desafio/1-temperatura/internal/entity"
)

type ZipCodeInputDto struct {
	ZipCode string `json:"cep"`
}

type ZipCodeOutputDto struct {
	Temp_C float32 `json:"temp_C"`
	Temp_F float32 `json:"temp_F"`
	Temp_K float32 `json:"temp_K"`
}

type TemperatureByZipCode struct {
	CityAdapter    entity.CityAdapterInterface
	WeatherAdapter entity.WeatherAdapterInterface
}

func NewTemperatureByZipCode(CityAdapter entity.CityAdapterInterface, WeatherAdapter entity.WeatherAdapterInterface) *TemperatureByZipCode {
	return &TemperatureByZipCode{
		CityAdapter:    CityAdapter,
		WeatherAdapter: WeatherAdapter,
	}
}

func (temp *TemperatureByZipCode) Execute(input ZipCodeInputDto) (ZipCodeOutputDto, error) {

	zipcode, err := entity.NewZipCode(input.ZipCode)
	if err != nil {
		return ZipCodeOutputDto{}, err
	}

	ctx := context.Background()
	ctx, cancelCtx := context.WithTimeout(ctx, 10*time.Second)
	defer cancelCtx()

	city, err := temp.CityAdapter.FindCity(ctx, &zipcode)
	if err != nil {
		return ZipCodeOutputDto{}, err
	}

	if !city.Exist() {
		return ZipCodeOutputDto{}, errors.New("can not find zip code")
	}

	weather, err := temp.WeatherAdapter.FindWeather(ctx, city)
	if err != nil {
		return ZipCodeOutputDto{}, err
	}

	resul := ZipCodeOutputDto{
		Temp_C: weather.Temp_C,
		Temp_F: weather.GetFahrenheit(),
		Temp_K: weather.GetKelvin(),
	}

	return resul, nil
}
