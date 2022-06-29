package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Balance(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	arg, err := strconv.ParseFloat(args, 32)

	if err != nil {
		log.Println("wrong args", args)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("Current balance: %.2f", arg))
	c.bot.Send(msg)
}
