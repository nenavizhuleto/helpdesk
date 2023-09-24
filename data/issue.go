package data

import "github.com/google/uuid"

type Issue struct {
	ID          string
	ClientID    string
	Company     string `form:"company"`
	Department  string `form:"department"`
	Name        string `form:"name"`
	PhoneNumber string `form:"phonenumber"`
	InnerNumber string `form:"innernumber"`
	Description string `form:"description"`
	Status      string
}

func NewIssue(clientID string) *Issue {
	return &Issue{
		ID:       uuid.NewString(),
		ClientID: clientID,
	}
}

var Subs = NewSubscriber()

func StoreIssue(id string, i *Issue) error {
	D.Store(id, *i)
	Subs.Notify(id)
	return nil
}

func GetIssuesById(id string) (map[string]Issue, error) {
	issues := D.Get(id)
	return issues, nil
}
