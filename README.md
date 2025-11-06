# 🤖 Telegram Bot in Go (with Cobra & Telebot)

This project is an example of building a **Telegram bot in Go** using:
- 🧠 [Cobra](https://github.com/spf13/cobra) — for CLI command handling;
- 💬 [Telebot.v3](https://pkg.go.dev/gopkg.in/telebot.v3) — for Telegram Bot API integration;
- 🔐 [Godotenv](https://github.com/joho/godotenv) — for loading environment variables.

The bot demonstrates clean Go project architecture with:
- modular package structure (`cmd`, `internal/bot`, `internal/config`);
- environment-based configuration;
- message and command handlers.

---

## 🧰 Tech Stack

| Component | Purpose |
|------------|----------|
| **Go 1.23+** | Programming language |
| **Cobra** | Command-line interface management |
| **Telebot.v3** | Telegram Bot SDK |
| **Godotenv** | Environment variable management |
| **Docker (optional)** | Containerization support |

---

## 🚀 How to Run Locally

### 1️⃣ Clone the repository
```bash
git clone https://github.com/BeardedHunter666/telegram-bot-go.git
cd telegram-bot-go
```
---

### 2️⃣ Create `.env` file

In the project root, create a `.env` file and add your TELE_TOKEN:
```bash
TELE_TOKEN=123456789:ABCdefGhiJKlmNoPQRstuVWxyZ
```
🛑 **Do not commit `.env` to GitHub!**  
(It’s already included in `.gitignore`)

---

### 3️⃣ Install dependencies
```bash
go mod tidy
```
---

### 4️⃣ Run the bot
```bash
go run main.go
```

You should see in your terminal:
```bash
🚀 Starting Telegram Bot...
✅ Bot is running...
```

---

### 5️⃣ Test in Telegram

Open your bot in Telegram:
t.me/<your_bot_name>_bot

💬 Available Commands
| Command        | Description                  |
| -------------- | ---------------------------- |
| `/start`       | Greeting and short intro     |
| `/help`        | List of available commands   |
| `/echo <text>` | Echo back user message       |
| Any text       | The bot repeats your message |

---

## 🧩 Project Structure
```bash
telegram-bot-go/
├── cmd/
│   └── root.go              # Cobra CLI entry point
├── internal/
│   ├── bot/
│   │   ├── bot.go           # Bot initialization
│   │   └── handlers.go      # Command handlers
│   └── config/
│       └── config.go        # Environment variable management
├── .env                     # Bot token (not tracked by Git)
├── .gitignore
├── go.mod
├── go.sum
└── main.go
```

---

## 🧠 Example Handlers
```go
b.Handle("/start", func(c tele.Context) error {
    return c.Send("👋 Hi there! I'm a Telegram bot written in Go. Type something!")
})

b.Handle("/help", func(c tele.Context) error {
    return c.Send("🤖 Commands:\n/start - greeting\n/help - list of commands\n/echo <text> - repeats your message")
})

b.Handle("/echo", func(c tele.Context) error {
    if c.Message().Payload == "" {
        return c.Send("⚠️ Usage: /echo <text>")
    }
    return c.Send("🪞 " + c.Message().Payload)
})
```

---

## 🧪 Usage Example
```bash
User: /start
Bot: 👋 Hi there! I'm a Telegram bot written in Go.

User: /echo Hello world!
Bot: 🪞 Hello world!
```

---