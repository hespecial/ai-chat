package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	LLM struct {
		BaseUrl string
		ApiKey  string
	}
	WeatherApi struct {
		Url  string
		Key  string
		City string
	}
	HotNewsApi struct {
		Url string
	}
}
