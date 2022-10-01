package main

import (
	"log"
	"os"

	"github.com/ftikhonin/BuddyBudgetBot/internal/app/commands"
	dbmanager "github.com/ftikhonin/BuddyBudgetBot/internal/app/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panicf("Authorization failed. Error: %s", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	bot.Debug = true

	dbmanager.InitDB()

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	commander := commands.NewCommander(bot)
	// var inlineArg bool
	var command string
	// TODO: add a goroutine
	for update := range updates {

		if update.Message == nil { // If we got a message
			continue
		}

		args := update.Message.CommandArguments()
		if args != "" {
			// inlineArg = true
			switch update.Message.Command() {
			case "help":
				commander.Help(update.Message)
			case "income":
				commander.IncomeInline(update.Message)
			case "expense":
				commander.Expense(update.Message)
			case "balance":
				commander.Balance(update.Message)
			case "list":
				commander.List(update.Message)
			case "register":
				commander.Register(update.Message)
			default:
				commander.Default(update.Message)
			}
			command = ""
		} else {
			if command == "" {
				command = update.Message.Command()
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter an "+command+".")
				bot.Send(msg)
			} else {
				switch command {
				// case "help":
				// 	commander.Help(update.Message)
				case "income":
					commander.Income(update.Message)
					// case "expense":
					// 	commander.Expense(update.Message)
					// case "balance":
					// 	commander.Balance(update.Message)
					// case "list":
					// 	commander.List(update.Message)
					// case "register":
					// 	commander.Register(update.Message)
					// default:
					// 	commander.Default(update.Message)
				}

			}

			// inlineArg = false

		}

		// if !inlineArg && command != "" {
		// 	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter an "+command+".")
		// 	bot.Send(msg)
		// }

	}
}
