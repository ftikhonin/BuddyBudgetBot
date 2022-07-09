package commands

import (
	dbmanager "github.com/ftikhonin/BuddyBudgetBot/internal/app/db"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) AddAccount(inputMessage *tgbotapi.Message) {
	dbmanager.AddAccount(inputMessage.Chat.ID, 300)
}
