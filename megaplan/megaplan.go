package megaplan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"application/auth"
	"application/models"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

type MegaPlan struct {
	Url   string
	Token Token
}

var MP *MegaPlan

type Employee struct {
	ID string `json:"id"`
}

type TaskDTO struct {
	Name        string   `json:"name"`
	Subject     string   `json:"subject"`
	Responsible Employee `json:"responsible"`
	IsUrgent    bool     `json:"isUrgent"`
	IsTemplate  bool     `json:"isTemplate"`
}

type TaskEvent struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Status  string `json:"status"`
}

type AuthOpt struct {
	Username       string
	Password       string
	GrantType      string
	AccessTokenUrl string
}

func NewAuthOpt(username, password string) *AuthOpt {
	return &AuthOpt{
		Username:       username,
		Password:       password,
		GrantType:      "password",
		AccessTokenUrl: "/auth/access_token",
	}
}

func (mp *MegaPlan) HandleCreateTask(i *auth.Identity, t *models.Task) (*models.Task, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	task_name := fmt.Sprintf("%s", t.Name)
	task_subject := fmt.Sprintf(`
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
    `,
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

	var responsible struct {
		ID string `json:"id"`
	}
	var task struct {
		Name        string      `json:"name"`
		Subject     string      `json:"subject"`
		Responsible interface{} `json:"responsible"`
		IsUrgent    bool        `json:"isUrgent"`
		IsTemplate  bool        `json:"isTemplate"`
	}

	task.Name = task_name
	task.Subject = task_subject
	task.Responsible = responsible
	task.IsUrgent = false
	task.IsTemplate = false

	jsonData, err := json.Marshal(task)
	log.Printf("JSON: %v", string(jsonData))
	log.Printf("URL: %v", string(mp.Url+"/task"))
	req, err := http.NewRequest("POST", mp.Url+"/task", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header["AUTHORIZATION"] = []string{mp.getToken()}
	res, err := client.Do(req)
	if res.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", res.StatusCode)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	log.Println(string(body))
	var response struct {
		Meta Meta        `json:"meta"`
		Data models.Task `json:"data"`
	}
	json.Unmarshal(body, &response)

	return &response.Data, nil
}

func setField(w *multipart.Writer, fieldName string, value string) error {
	fw, err := w.CreateFormField(fieldName)
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, strings.NewReader(value))
	if err != nil {
		return err
	}

	return nil
}

func (mp *MegaPlan) MustAuthenticateWithPassword(auth *AuthOpt) *MegaPlan {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)
	if err := setField(w, "username", auth.Username); err != nil {
		log.Fatalf("Error while authenticating megaplan: %v", err)
	}
	if err := setField(w, "password", auth.Password); err != nil {
		log.Fatalf("Error while authenticating megaplan: %v", err)
	}
	if err := setField(w, "grant_type", auth.GrantType); err != nil {
		log.Fatalf("Error while authenticating megaplan: %v", err)
	}
	w.Close()
	req, err := http.NewRequest("POST", mp.Url+auth.AccessTokenUrl, bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Fatalf("Error while authenticating megaplan: %v", err)
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	res, err := client.Do(req)
	if res.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", res.StatusCode)
		log.Fatalf("Error while authenticating megaplan: %v", err)
	}
	token := Token{}
	json.NewDecoder(res.Body).Decode(&token)
	mp.Token = token
	return mp
}

func New(url string, authOpt *AuthOpt) *MegaPlan {
	mp := &MegaPlan{
		Url: url,
	}

	return mp
}

type Pagination struct {
	Count int `json:"count"`
	Limit int `json:"limit"`
}

type Meta struct {
	Status     int        `json:"status"`
	Errors     []string   `json:"errors"`
	Pagination Pagination `json:"pagination"`
}

type Response struct {
	Meta Meta          `json:"meta"`
	Data []interface{} `json:"data"`
}

func (mp *MegaPlan) getToken() string {
	return "Bearer " + mp.Token.AccessToken
}

func (mp *MegaPlan) doRequest(method, url string, body io.Reader) (Response, error) {
	req, _ := http.NewRequest(method, mp.Url+url, body)
	req.Header["AUTHORIZATION"] = []string{mp.getToken()}

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := client.Do(req)
	if res.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", res.StatusCode)
		return Response{}, err
	}
	rsp := Response{}
	json.NewDecoder(res.Body).Decode(&rsp)
	return rsp, nil
}

func (mp *MegaPlan) Get(url string) (Response, error) {
	return mp.doRequest("GET", url, nil)
}

func (mp *MegaPlan) Post(url string, data interface{}) (Response, error) {
	return mp.doRequest("GET", url, nil)
}
