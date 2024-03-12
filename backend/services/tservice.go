package services

import (
	"back/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var apiURL string = "https://api.openai.com/v1/chat/completions"

func GenerateText(message string) (string, error) {

	godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	messages := make([]models.Message, 0)

	messages = append(messages, models.Message{Role: "user", Content: message})

	body := models.Request{Model: "gpt-4", Messages: messages}

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

	responseBody, err := io.ReadAll(res.Body)

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

func HandleGenerateText(c *gin.Context) {
	var requestData map[string]string

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	output, err := GenerateText(requestData["message"])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating text"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": output})
}
