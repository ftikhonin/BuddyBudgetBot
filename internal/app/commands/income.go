package commands

import (
	"fmt"
	"log"
	"strconv"

	dbmanager "github.com/ftikhonin/BuddyBudgetBot/internal/app/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) IncomeCore(ID int64, args string) {

	arg, err := strconv.ParseFloat(args, 32)

	if err != nil {
		log.Println("wrong args", args)
	}

	if arg <= 0 {

		msg := tgbotapi.NewMessage(ID, "Please, enter a positive number")
		c.bot.Send(msg)

	} else {

		balance, err := dbmanager.SetIncome(ID, arg, "test")
		if err != nil {
			log.Println("something wrong")
		}

		if s, err := strconv.ParseFloat(balance, 32); err == nil {

			msg := tgbotapi.NewMessage(ID, fmt.Sprintf("Current balance: %.2f", s))
			c.bot.Send(msg)

		}
	}

}

func (c *Commander) IncomeInline(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	c.IncomeCore(inputMessage.Chat.ID, args)

}

func (c *Commander) Income(inputMessage *tgbotapi.Message) {
	c.IncomeCore(inputMessage.Chat.ID, inputMessage.Text)
}
