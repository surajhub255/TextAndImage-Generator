package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"errors"
	"back/models"
	
)
var apiURL string = "https://api.openai.com/v1/chat/completions"

var apiKey string = "Your api  key"



func GenerateText(message string) (string, error) {
	messages := make([]models.Message, 0)

	messages = append(messages, models.Message{Role: "user", Content: message})

	body := models.Request{Model: "gpt-3.5-turbo", Messages: messages}

	bodyBytes, _ := json.Marshal(body)

	r, _ := http.NewRequest("POST", apiURL, bytes.NewReader(bodyBytes))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return "", errors.New("could not generate text")
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", errors.New("could not generate text")
	}

	output := models.Response{}

	err = json.Unmarshal(responseBody, &output)
	if err != nil {
		return "", errors.New("could not generate text")
	}

	if len(output.Choices) == 0 {
		return "", errors.New("could not generate text")
	}

	return output.Choices[0].Message.Content, nil
}

func HandleGenerateText(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	output, err := GenerateText(requestData["message"])
	if err != nil {
		http.Error(w, "Error generating text", http.StatusInternalServerError)
		return
	}

	response := map[string]string{"result": output}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
