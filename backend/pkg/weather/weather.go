package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiConfig struct {
	Url  string
	Key  string
	City string
}

func New(url, key, city string) *ApiConfig {
	return &ApiConfig{
		Url:  url,
		Key:  key,
		City: city,
	}
}

type Response struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	Lives    []struct {
		Province         string `json:"province"`
		City             string `json:"city"`
		Adcode           string `json:"adcode"`
		Weather          string `json:"weather"`
		Temperature      string `json:"temperature"`
		Winddirection    string `json:"winddirection"`
		Windpower        string `json:"windpower"`
		Humidity         string `json:"humidity"`
		Reporttime       string `json:"reporttime"`
		TemperatureFloat string `json:"temperature_float"`
		HumidityFloat    string `json:"humidity_float"`
	} `json:"lives"`
}

func (c *ApiConfig) Query() (*Response, error) {
	resp, err := http.Get(fmt.Sprintf("%s?key=%s&&city=%s", c.Url, c.Key, c.City))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r Response
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
