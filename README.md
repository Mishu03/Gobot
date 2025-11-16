
---

# 🧠 Discord AI Chatbot (Go + OpenAI)

A lightweight and fast AI-powered Discord chatbot built using **Golang**, **DiscordGo**, and the **OpenAI API**.

This bot listens to Discord commands and responds with:

* Latency checks
* Greetings
* AI-generated answers using the `!ask` command

---

## ⭐ Features

### 🔹 AI Chat Command

Ask the bot anything using:

```
!ask What is Golang?
```

The bot sends your question to OpenAI and replies with an intelligent answer.

### 🔹 Basic Commands

```
!ping   → Pong!
!hello  → Simple greeting
```

### 🔹 Built in Go

Fast, concurrent, simple to deploy, and easy to extend.

---

## 📦 Tech Stack

* **Go (Golang)**
* **DiscordGo** – Discord API wrapper
* **OpenAI API**
* **Environment variables** for secure key management

---

## 🚀 Getting Started

### 1. Clone the Repository

```
git clone https://github.com/yourname/discord-go-chatbot.git
cd discord-go-chatbot
```

### 2. Set Environment Variables

```
setx DISCORD_BOT_TOKEN "your-bot-token"
setx OPENAI_API_KEY "your-openai-key"
```

> Restart your terminal after setting them.

### 3. Run the Bot

```
go run main.go
```

---

## 📁 Project Structure

```
/discord-bot
│── main.go
│── go.mod
│── README.md
```

---

## 🎯 Bot Commands

| Command           | Description               |
| ----------------- | ------------------------- |
| `!ping`           | Check if the bot is alive |
| `!hello`          | Simple greeting           |
| `!ask <question>` | Ask the AI anything       |

---

## 🤝 Contributions

Contributions, issues, and feature requests are welcome.
Feel free to submit a PR!

---

## 📄 License

This project is licensed under the **MIT License**.

---
e!
