package commands

import (
	"fmt"
	"log"
	"strconv"

	dbmanager "github.com/ftikhonin/BuddyBudgetBot/internal/app/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Income(inputMessage *tgbotapi.Message) (inlineArg bool) {

	args := inputMessage.CommandArguments()
	if args == "" {
		return false
	}

	arg, err := strconv.ParseFloat(args, 32)

	if err != nil {
		log.Println("wrong args", args)
		return true
	}

	if arg <= 0 {

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Please, enter a positive number")
		c.bot.Send(msg)

	} else {

		balance, err := dbmanager.SetIncome(inputMessage.Chat.ID, arg, "test")
		if err != nil {
			log.Println("something wrong")
			return true

		}

		if s, err := strconv.ParseFloat(balance, 32); err == nil {

			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Current balance: %.2f", s))
			c.bot.Send(msg)

		}
	}
	return true

}
