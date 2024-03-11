package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"back/models"
)

var apiURLI string = "https://api.openai.com/v1/images/generations"

var apiKeyI string = "Your api key"


func GenerateImage(w http.ResponseWriter, req *http.Request) {

	text := req.URL.Query().Get("text")

	body := models.RequestI{Prompt: text, N: 1, ResponseFormat: "url", Size:models.Low}

	bodyBytes, _ := json.Marshal(body)

	r, _ := http.NewRequest("POST", apiURLI, bytes.NewReader(bodyBytes))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+apiKeyI)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	output :=models.ResponseI{}

	err = json.Unmarshal(responseBody, &output)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output.Data[0].Url)
}
