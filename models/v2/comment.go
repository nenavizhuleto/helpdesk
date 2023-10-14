package models

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	User    User   `json:"user"`
}
