package main

import (
	"flag"
	"fmt"

	"weatherProject/geo"
	"weatherProject/weather"
)

func main() {
	fmt.Println("MY WEATHER PROJECT")
	city := flag.String("city", "Tashkent", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()

	fmt.Println(*city)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Printf("Запрос вернул ошибку: %v\n", err.Error())
	}
	fmt.Println(geoData)

	weatherData := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
