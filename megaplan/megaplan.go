package megaplan

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

type MegaPlan struct {
	Url         string
	Responsible string
	Token       Token
}

var MP *MegaPlan

type AuthOpt struct {
	Username       string
	Password       string
	GrantType      string
	AccessTokenUrl string
	Responsible    string
}

func NewAuthOpt(username, password, responsible string) *AuthOpt {
	return &AuthOpt{
		Username:       username,
		Password:       password,
		GrantType:      "password",
		AccessTokenUrl: "/auth/access_token",
		Responsible:    responsible,
	}
}

func (mp *MegaPlan) doRequest(method, url string, body interface{}, response interface{}) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	jsonData, err := json.Marshal(body)
	req, err := http.NewRequest(method, mp.Url+url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header["AUTHORIZATION"] = []string{mp.getToken()}
	res, err := client.Do(req)
	if res.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", res.StatusCode)
		bytes, _ := io.ReadAll(res.Body)
		log.Printf("Body: %s", string(bytes))
		return errors.New("Status Not OK")
	}
	defer res.Body.Close()

	bytes, _ := io.ReadAll(res.Body)
	json.Unmarshal(bytes, &response)
	return nil
}

func (mp *MegaPlan) Get(url string) *Response {
	var res Response
	if err := mp.doRequest("GET", url, nil, &res); err != nil {
		panic(err)
	}

	log.Printf("F: %v", res)

	return &res
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
		Url:         url,
		Responsible: authOpt.Responsible,
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
