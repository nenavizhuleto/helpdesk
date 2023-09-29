package models

import (
	"database/sql"

	"github.com/google/uuid"

	"application/data"
)

type User struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Phone string `json:"phone" db:"phone"`
}

func GetUserFromDevice(device *Device) (*User, error) {
	db := data.DB
	var user User
	if device.UserID.Valid {
		if err := db.Get(&user, "SELECT * FROM users WHERE id = $1", device.UserID.String); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, sql.ErrNoRows
}

func NewUser(ip, name, phone string) error {
	db := data.DB

	user := User{
		ID:    uuid.NewString(),
		Name:  name,
		Phone: phone,
	}

	_, err := db.NamedExec("INSERT INTO users VALUES (:id, :name, :phone)", &user)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE devices SET user_id = $1 WHERE ip = $2", user.ID, ip)
	if err != nil {
		return err
	}

	return nil
}
