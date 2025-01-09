package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// sendTelegramMessage sends a message to a Telegram chat
func sendTelegramMessage(msg string, chatId string) error {
	// If the chat ID is empty, use the default chat ID from the configuration which is send to System Admin
	if chatId == "" {
		chatId = AppConfig.ChatId
	} 
	// Get bot token from the configuration
	botToken := AppConfig.BotToken

	// Construct the URL for the Telegram API
	url := "https://api.telegram.org/bot" + botToken + "/sendMessage"

	// Create the payload with the chat ID, message text and parse mode
	payload := map[string]string{
		"chat_id":    chatId,
		"text":       msg,
		"parse_mode": "Markdown",
	}

	// Convert the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Send the POST request to the Telegram API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

    // Check if the request was successful
    if resp.StatusCode != http.StatusOK {
        bodyBytes, _ := io.ReadAll(resp.Body)
        bodyString := string(bodyBytes)
        return fmt.Errorf("failed to send message, status code: %d, reason: %s", resp.StatusCode, bodyString)
    }

    return nil
}

// sendMessageToUser sends a message to a specific user
func sendMessageToUser(w http.ResponseWriter, message, userID string) {
	err := sendTelegramMessage(message, userID)
	if err != nil {
		log.Printf("Failed to get telegram ID: %v", err)
	} else {
		msgLog := fmt.Sprintf("Message sent to Telegram userID %s", userID)
		SendResponse(w, msgLog)
	}
}