package commands

import (
	"fmt"
	"log"
	"strconv"

	dbmanager "github.com/ftikhonin/BuddyBudgetBot/internal/app/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Expense(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	arg, err := strconv.ParseFloat(args, 32)

	if err != nil {

		log.Println("wrong args", args)
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Please enter the correct amount"))
		c.bot.Send(msg)
		return
	}

	if arg <= 0 {

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Please enter a positive number"))
		c.bot.Send(msg)

	} else {

		balance, err := dbmanager.SetExpense(inputMessage.Chat.ID, arg, "test")

		if err != nil {
			log.Println("something wrong:" + err.Error())
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, err.Error())
			c.bot.Send(msg)
			return
		}

		if s, err := strconv.ParseFloat(balance, 32); err == nil {
			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Current balance: %.2f", s))
			c.bot.Send(msg)
		}
	}
}
