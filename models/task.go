package models

import (
	"errors"
	"log"
	"time"

	"application/data"
)

type TaskTimeCreated struct {
	Value string `json:"value"`
}

func (tc *TaskTimeCreated) Scan(src interface{}) error {
	var source []byte
	switch src := src.(type) {
	case string:
		source = []byte(src)
	case []byte:
		source = src
	default:
		return errors.New("Incompatible type for TaskTimeCreated")
	}
	log.Println("datetime: ", string(source))
	format := "2006-01-02 15:04:05Z07:00"
	date, err := time.Parse(format, string(source))
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
