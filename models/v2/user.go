package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"application/data"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (u *User) Create(ip string) error {
	db := data.DB

	_, err := db.NamedExec("INSERT INTO users VALUES (:id, :name, :phone)", u)
	if err != nil {
		return err
	}

	coll := data.GetCollection(devices)
	if err := coll.FindOneAndUpdate(nil, bson.D{{Key: "ip", Value: ip}}, bson.D{{Key: "$set", Value: bson.M{"user": u}}}).Err(); err != nil {
		return fmt.Errorf("user: %w", err)
	}

	return nil
}
