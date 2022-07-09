package commands

import (
	"fmt"
	"log"
	"strconv"

	dbmanager "github.com/ftikhonin/BuddyBudgetBot/internal/app/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Balance(inputMessage *tgbotapi.Message) {

	balance, err := dbmanager.GetBalance(inputMessage.Chat.ID)
	if err != nil {
		log.Println("something wrong")
		return
	}

	if s, err := strconv.ParseFloat(balance, 32); err == nil {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Current balance: %.2f", s))
		c.bot.Send(msg)
	}
}
