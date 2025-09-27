package news

import (
	"encoding/json"
	"io"
	"net/http"
)

type ApiConfig struct {
	Url string
}

func New(url string) *ApiConfig {
	return &ApiConfig{
		Url: url,
	}
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Desc  string `json:"desc"`
		Hot   string `json:"hot"`
		Img   string `json:"img"`
		Index int    `json:"index"`
		Title string `json:"title"`
		Url   string `json:"url"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

func (c *ApiConfig) Query() (*Response, error) {
	resp, err := http.Get(c.Url)
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
