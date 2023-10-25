package telegram

import "helpdesk/internals/models/v3/user"


func (t *TelegramNotificator) NotifyUser(user *user.TelegramUser) {
	chat := t.chats.Get(user.ChatID)
	authorized, ok := chat.Data["authorized"]
	if !ok {
		return
	}

	if authorized.(bool) {
		t.SendMessageDirectlyToChat(user.ChatID, "Вам новое уведомление от Helpdesk")
	}


}
