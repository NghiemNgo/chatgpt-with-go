package handler

import (
	"go.tienngay/chatGpt/database"
	"github.com/gofiber/fiber/v2"
	"go.tienngay/pkg/mongo/repositories/chatGpt"
	MongoEntities "go.tienngay/pkg/mongo/entities"
	"fmt"
	"net/http"
	gogpt "github.com/sashabaranov/go-gpt3"
	"go.tienngay/chatGpt/config"
	"context"
	"time"
)

type ChatInput struct {
    Prompt string
    Email string
}


func ChatGptService() chatGpt.ChatGptService {
	db := database.MG
	gptCollection := db.Collection("chat_gpt")
	gptRepo := chatGpt.NewRepo(gptCollection)
	gptService := chatGpt.NewService(gptRepo)
	return gptService;
}

func Chat(c *fiber.Ctx) error {
	var gptService = ChatGptService()
	// Declare a new ChatInput struct.
	var chatInput ChatInput

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    if err := c.BodyParser(&chatInput); err != nil {
        return c.JSON(fiber.Map{"status": http.StatusBadRequest, "message": "Không phân giải được dữ liệu.", "data": nil})
    }
    if (chatInput.Prompt == "") {
    	return c.JSON(fiber.Map{"status": http.StatusBadRequest, "message": "Câu hỏi không được để trống", "data": nil})
    }

    gpt := gogpt.NewClient(config.Config("CHAT_GPT_KEY"))
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 4000,
		Prompt:    chatInput.Prompt,
	}
	resp, err := gpt.CreateCompletion(ctx, req)
	if err != nil {
		return c.JSON(fiber.Map{"status": http.StatusBadRequest, "message": "ChatGPT không thể trả lời câu hỏi", "data": nil})
	}

	var chat MongoEntities.ChatGpt
	chat.Prompt = chatInput.Prompt
	chat.Output = resp.Choices[0].Text
	chat.CreatedBy = chatInput.Email
	chat.CreatedAt = time.Now().Unix()
	insertId := gptService.Insert(chat)
	fmt.Println(insertId)

	return c.JSON(fiber.Map{"status": http.StatusOK, "message": "Success", "data": resp.Choices[0].Text})
}
