package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"weatherProject/geo"
)

var (
	ErrWrongFormat = errors.New("ERROR_FORMAT")
	ErrToParse     = errors.New("ERROR_URL")
	ErrToRequest   = errors.New("ERROR_HTTP")
	ErrToRead      = errors.New("ERROR_READ_BODY")
)

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(ErrToParse)
		return "", err
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(ErrToRequest)
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(ErrToRead)
		return "", err
	}
	return string(body), nil
}
