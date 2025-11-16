package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/bwmarrin/discordgo"
)

// -------------------------
// AI FUNCTION
// -------------------------
func askAI(prompt string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{Role: "user", Content: prompt},
			},
		},
	)

	if err != nil {
		return "Error: " + err.Error()
	}

	return resp.Choices[0].Message.Content
}

// -------------------------
// MAIN FUNCTION
// -------------------------
func main() {
	fmt.Println("OPENAI_API_KEY:", os.Getenv("OPENAI_API_KEY"))

	token := os.Getenv("DISCORD_BOT_TOKEN")

	if token == "" {
		fmt.Println("DISCORD_BOT_TOKEN not set")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	dg.AddHandler(messageHandler)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is running... Press CTRL+C to exit")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	<-sc

	dg.Close()
}

// -------------------------
// MESSAGE HANDLER
// -------------------------
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	// COMMAND: !ping
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// COMMAND: !hello
	if m.Content == "!hello" {
		s.ChannelMessageSend(m.ChannelID, "Hello! How can I help?")
	}

	// COMMAND: !ask <question>
	if strings.HasPrefix(m.Content, "!ask ") {
		question := strings.TrimPrefix(m.Content, "!ask ")
		answer := askAI(question)
		s.ChannelMessageSend(m.ChannelID, answer)
	}
}
