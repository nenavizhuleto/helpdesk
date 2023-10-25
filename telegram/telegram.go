package telegram

import (
	"fmt"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandHandlerFunc func(tg.Update)

type TelegramNotificator struct {
	b        *tg.BotAPI
	u        tg.UpdateConfig
	commands map[string]CommandHandlerFunc
	chats    Chats
}

func NewTelegramNotificator(token string) (*TelegramNotificator, error) {
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("telegram: %w", err)
	}

	u := tg.NewUpdate(0)
	u.Timeout = 60

	return &TelegramNotificator{
		b:        bot,
		u:        u,
		commands: make(map[string]CommandHandlerFunc),
		chats:    make(Chats),
	}, nil
}

func (t *TelegramNotificator) SetDebug(v bool) {
	t.b.Debug = v
}

func (t *TelegramNotificator) AddCommandHandler(command string, handler CommandHandlerFunc) {
	t.commands[command] = handler
}

func (t *TelegramNotificator) Run() {
	t.AddCommandHandler("start", t.StartCommand)
	t.AddCommandHandler("help", t.HelpCommand)
	t.AddCommandHandler("login", t.LoginCommand)
	t.AddCommandHandler("info", t.InfoCommand)

	log.Printf("Starting telegram bot...")

	updates := t.b.GetUpdatesChan(t.u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			t.HandleCommand(update)
		} else {
			t.HandleMessage(update)
		}
	}
}
