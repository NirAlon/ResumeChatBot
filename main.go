package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

func main() {
	resumeBytes, err := ioutil.ReadFile("resume.txt")
	if err != nil {
		panic("Could not read resume.txt: " + err.Error())
	}
	resumeText := string(resumeBytes)

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY enviroment variable not set")
	}
	client := openai.NewClient(apiKey)

	router := gin.Default()
	// router.Use(corsMiddleware()) // uncoment this for local tests

	router.POST("/chat", func(ctx *gin.Context) {
		var req ChatRequest
		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a helpful assistant that answers reqruters and tech people questions about nir alon:\n\n" + resumeText,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: req.Message,
			},
		}
		resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Messages:    messages,
			MaxTokens:   150,
			Temperature: 0.2,
		})
		if err != nil {
			fmt.Println("OpenAI error:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"response": "Error contacting OpenAI"})
			return
		}
		if len(resp.Choices) == 0 {
			fmt.Println("OpenAI response had 0 choices! Full response:", resp)
			ctx.JSON(http.StatusInternalServerError, gin.H{"response": "No response from OpenAI"})
			return
		}

		answer := resp.Choices[0].Message.Content
		ctx.JSON(http.StatusOK, ChatResponse{Response: answer})

	})

	router.Run(":8080")

}


// func corsMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }
//         c.Next()
//     }
// }
