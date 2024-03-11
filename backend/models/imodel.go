package models

import (

)


type RequestI struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
}

const (
	Low    string = "256x256"
	medium string = "512x512"
	high   string = "1024x1024"
)

type ResponseI struct {
	Created int `json:"created"`
	Data    []struct {
		Url string `json:"url"`
	} `json:"data"`
}