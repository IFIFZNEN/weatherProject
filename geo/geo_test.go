package geo_test

import (
	"errors"
	"testing"

	"weatherProject/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arange -подготовка, expected результат, данные для функции
	city := "Tashkent"
	expected := geo.GeoData{
		City: "Tashkent",
	}
	// Act -выполнение функции
	got, err := geo.GetMyLocation(city)
	// Assert - сравнение результата с expected
	if err != nil {
		t.Error(err.Error())
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "QWERTY"
	_, err := geo.GetMyLocation(city)
	if !errors.Is(err, geo.ErrNoCity) {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrNoCity, err)
	}
}
