package models

import (
	"errors"
	"strings"
	"time"

	"application/data"
)

type TaskTimeCreated struct {
	Value string `json:"value"`
}

var TaskTimeFormatString = "2006-01-02 15:04:05Z07:00"

func (tc *TaskTimeCreated) Scan(src interface{}) error {
	source, ok := src.(string)
	if !ok {
		return errors.New("Incompitable datetime type")
	}
	// FIXME: 'T' symbol occasionally returned in timeCreated from megaplan
	dt := strings.Replace(string(source), "T", " ", 1)
	date, err := time.Parse(TaskTimeFormatString, dt)
	if err != nil {
		return errors.New("Couldn't parse datetime")
	}
	tc.Value = date.Format(time.DateTime)
	return nil
}

type TaskFilter struct {
	ID string `json:"id"`
}

type Task struct {
	ID          string          `json:"id" db:"id"`
	HumanNumber int             `json:"humanNumber" db:"human_number"`
	UserID      string          `json:"user_id" db:"user_id"`
	Name        string          `json:"name" db:"name"`
	Subject     string          `json:"subject" db:"subject"`
	Status      string          `json:"status" db:"status"`
	TimeCreated TaskTimeCreated `json:"timeCreated" db:"time_created"`
}

func (t *Task) Prettify() error {
	oldTime := t.TimeCreated.Value
	oldTime = strings.Replace(oldTime, "T", " ", 1)
	tm, err := time.Parse(TaskTimeFormatString, oldTime)
	if err != nil {
		return err
	}
	t.TimeCreated.Value = tm.Format(time.DateTime)
	return nil
}

func GetTasksForUser(uid string) ([]Task, error) {
	db := data.DB

	var tasks []Task
	if err := db.Select(&tasks, "SELECT * FROM tasks WHERE user_id = ?", uid); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTaskForUser(uid string, task *Task) error {
	db := data.DB
	task.UserID = uid
	_, err := db.NamedExec("INSERT INTO tasks VALUES (:id, :human_number, :user_id, :name, :subject, :status, :time_created.value)", task)
	if err != nil {
		return err
	}
	return nil
}
