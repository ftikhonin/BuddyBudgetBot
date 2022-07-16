package commands

import (
	"log"
	"time"

	dbmanager "github.com/ftikhonin/BuddyBudgetBot/internal/app/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	if args == "" {
		log.Println("wrong args", args)
		return
	}
	var forDate string
	if args == "today" {
		forDate = time.Now().Format("20060102")
	}

	if args == "month" {
		forDate = time.Now().AddDate(0, -1, 0).Format("20060102")
	}

	if forDate != "" {
		balance, _ := dbmanager.GetList(inputMessage.Chat.ID, forDate)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, balance)

		c.bot.Send(msg)
	}

}
