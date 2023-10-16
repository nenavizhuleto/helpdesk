package megaplan

import "time"

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Name      string `json:"name,omitempty"`
}

type TimeCreated struct {
	Value time.Time `json:"value,omitempty"`
}

type Comment struct {
	Content string `json:"content"`
}

// Piece of information to create task in Megaplan
type TaskDTO struct {
	ID            string       `json:"id,omitempty"`
	Name          string       `json:"name"`
	Subject       string       `json:"subject"`
	Responsible   Employee     `json:"responsible"`
	IsUrgent      bool         `json:"isUrgent"`
	IsTemplate    bool         `json:"isTemplate"`
	TimeCreated   *TimeCreated `json:"timeCreated,omitempty"`
	Activity      *TimeCreated `json:"activity,omitempty"`
	Status        string       `json:"status,omitempty"`
	LastComment   *Comment     `json:"lastComment,omitempty"`
	Comments      []Comment    `json:"comments,omitempty"`
	CommentsCount int          `json:"commentsCount,omitempty"`
}

type TaskEvent struct {
	Data  TaskDTO `json:"data"`
	Event string  `json:"event"`
	Model string  `json:"model"`
}
