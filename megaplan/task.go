package megaplan

import (
	"fmt"

	"application/auth/v1"
	"application/models/v1"
)

type TaskDTO struct {
	Name        string          `json:"name"`
	Subject     string          `json:"subject"`
	Responsible models.Employee `json:"responsible"`
	IsUrgent    bool            `json:"isUrgent"`
	IsTemplate  bool            `json:"isTemplate"`
}

type TaskEvent struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Status  string `json:"status"`
}

var TaskSubjectFormat = `
    <h2>%s от %s:</h2>
    <h3>Суть обращения:</h3>
    <p>%s</p>
    <hr/>
    <h3>Дополнительная информания:</h3>
    <ul>
    <li>Контакты: %s</li>
    <li>Устройство: %s (%s)</li>
    <li>Отдел: <br/>Название: %s <br/>Описание: %s <br/>Адрес: %s <br/>Контакты: %s</li>
    </ul>
    `

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

func (mp *MegaPlan) HandleCreateTask(i *auth.Identity, t *models.Task) (*models.Task, error) {
	task_name := fmt.Sprintf("%s", t.Name)
	task_subject := fmt.Sprintf(TaskSubjectFormat,
		i.Company.Name,
		i.User.Name,
		t.Subject,
		i.User.Phone,
		i.Device.IP,
		i.Subnet.Network,
		i.Branch.Name,
		i.Branch.Description,
		i.Branch.Address,
		i.Branch.Contacts,
	)

	responsible := models.Employee{
		ID: mp.Responsible,
	}

	task := models.Task{
		Name:        task_name,
		Subject:     task_subject,
		Responsible: responsible,
		IsUrgent:    false,
		IsTemplate:  false,
	}

	var response struct {
		Meta Meta        `json:"meta"`
		Data models.Task `json:"data"`
	}

	if err := mp.doRequest("POST", "/task", task, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
