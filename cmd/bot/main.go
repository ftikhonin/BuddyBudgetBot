package main

import (
	"log"
	"os"

	"github.com/ftikhonin/BuddyBudgetBot/internal/app/commands"
	dbmanager "github.com/ftikhonin/BuddyBudgetBot/internal/app/db"
	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	dbmanager.InitDB()

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	commander := commands.NewCommander(bot)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}

		switch update.Message.Command() {
		case "help":
			commander.Help(update.Message)
		case "income":
			commander.Income(update.Message)
		case "expense":
			commander.Expense(update.Message)
		case "balance":
			commander.Balance(update.Message)
		case "reg":
			commander.AddAccount(update.Message)
		default:
			commander.Default(update.Message)
		}

	}
}
