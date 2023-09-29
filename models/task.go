package models

type Task struct {
	ID      string `json:"id" db:"id"`
	UserID  string `json:"user_id" db:"user_id"`
	Name    string `json:"name" db:"name"`
	Subject string `json:"subject" db:"subject"`
	Status  string `json:"status" db:"status"`
}
