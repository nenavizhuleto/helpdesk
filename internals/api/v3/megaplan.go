package api

import (
	"encoding/json"
	"helpdesk/internals/megaplan"
	tk "helpdesk/internals/models/v3/task"
	"helpdesk/telegram"
	"log"

	"github.com/gofiber/fiber/v2"
)


func HandleMegaplanEvent(c *fiber.Ctx) error {
	var event megaplan.TaskEvent
	if err := c.BodyParser(&event); err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	if event.Model != "Task" {
		return c.SendStatus(200)
	}

	str, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("Event: %s", string(str))

	dto := event.Data
	task, err := tk.Get(dto.ID)
	if err != nil {
		log.Printf("event: %s", err.Error())
		return c.SendStatus(200)
	}

	needUpdate := false
	update := make(tk.UpdateEvent, 0)
	newStatus := dto.GetStatus()
	if task.Status != newStatus {
		task.Status = newStatus
		update = append(update, tk.StatusUpdate)
		needUpdate = true
	}
	newComment := dto.GetLastComment()
	if newComment != nil {
		task.Comments = append(task.Comments, *newComment)
		update = append(update, tk.CommentUpdate)
		needUpdate = true
	}
	if dto.Activity != nil && needUpdate {
		task.LastActivity = dto.Activity.Value
		update = append(update, tk.ActivityUpdate)
	}

	if needUpdate {
		if err := task.Save(); err != nil {
			log.Println(err)
			return c.SendStatus(500)
		}

		if tg, _ := task.User.GetTelegram(); tg != nil {
			telegram.Bot.NotifyUser(tg, task, update)
		}
	}

	return c.SendStatus(200)
}
