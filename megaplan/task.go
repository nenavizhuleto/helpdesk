package megaplan

import (
	"fmt"

	"application/models/v1"
)

func (mp *MegaPlan) HandleFetchTaskUpdates(t *models.Task) (*models.Task, error) {
	var response struct {
		Meta Meta        `json:"meta"`
		Data models.Task `json:"data"`
	}
	if err := mp.doRequest("GET", fmt.Sprintf("/task/%s", t.ID), nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

func (mp *MegaPlan) CreateTask(name, subject string) (*TaskDTO, error) {

	responsible := Employee{
		ID: mp.Responsible,
	}

	task := TaskDTO{
		Name:        name,
		Subject:     subject,
		Responsible: responsible,
		IsUrgent:    false,
		IsTemplate:  false,
	}

	var response struct {
		Meta Meta        `json:"meta"`
		Data TaskDTO `json:"data"`
	}

	if err := mp.doRequest("POST", "/task", task, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
