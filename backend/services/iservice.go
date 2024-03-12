package services

import (
	"back/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var apiURLI string = "https://api.openai.com/v1/images/generations"

func GenerateImage(c *gin.Context) {

	godotenv.Load()

	apiKeyI := os.Getenv("API_KEY")
	text := c.Query("text")

	body := models.RequestI{Prompt: text, N: 1, ResponseFormat: "url", Size: models.Low, Model: "dall-e-2"}

	bodyBytes, _ := json.Marshal(body)

	r, _ := http.NewRequest("POST", apiURLI, bytes.NewReader(bodyBytes))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+apiKeyI)
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request to OpenAI API"})
		return
	}

	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response body"})
		return
	}

	output := models.ResponseI{}

	err = json.Unmarshal(responseBody, &output)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding OpenAI API response"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"imageSrc": output.Data[0].Url})
}
