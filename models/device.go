package models

import (
	"database/sql"

	"application/data"
)

type Device struct {
	IP     string         `json:"ip" db:"ip"`
	UserID sql.NullString `json:"user_id" db:"user_id"`
	Type   string         `json:"type" db:"type"`
}

func NewDevice(ip string) (*Device, error) {
	db := data.DB

	var device Device
	if err := db.Get(&device, "SELECT * FROM devices WHERE ip = $1", ip); err != nil {
		device = Device{
			IP:   ip,
			Type: "PC",
		}
		_, err := db.NamedExec("INSERT INTO devices(ip, type) VALUES (:ip, :type)", &device)
		if err != nil {
			return &device, err
		}
	}

	return &device, nil
}
