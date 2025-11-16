🧠 Discord AI Chatbot (Go + OpenAI)

A lightweight and fast AI-powered Discord chatbot built using Golang, DiscordGo, and the OpenAI API.

This bot listens to commands inside Discord and can respond with:

Latency checks

Greetings

AI-generated answers using the !ask command

⭐ Features
🔹 AI Chat Command

Ask the bot anything:

!ask What is Golang?


The bot sends your question to OpenAI and replies with an intelligent answer.

🔹 Basic Commands
!ping   → Pong!
!hello  → Friendly greeting

🔹 Fully written in Go

Fast, concurrent, and simple to deploy.

📦 Tech Stack

Go (Golang)

DiscordGo (Discord API wrapper)

OpenAI API

Environment variables for secure key management

🚀 Getting Started
1. Clone the repo
git clone https://github.com/yourname/discord-go-chatbot.git
cd discord-go-chatbot

2. Set environment variables
setx DISCORD_BOT_TOKEN "your-bot-token"
setx OPENAI_API_KEY "your-openai-key"

3. Run the bot
go run main.go

📁 Project Structure
/discord-bot
│── main.go
│── go.mod
│── README.md

🎯 Commands
Command	Description
!ping	Check if bot is alive
!hello	Simple greeting
!ask <question>	Ask AI any question
🤝 Contributions

Feel free to open issues or submit PRs!

📄 License

MIT License
