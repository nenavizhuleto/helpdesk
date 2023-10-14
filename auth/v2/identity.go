package auth

import (
	"fmt"

	"application/models/v2"
)

func GetIdentity(ip string) (*models.Device, error) {
	return models.DeviceGetByIP(ip)
}

func Register(ip string, user *models.User) (*models.Device, error) {
	dev, err := MakeIdentity(ip)
	if err != nil {
		return nil, fmt.Errorf("identity: %w", err)
	}
	if err := dev.Create(); err != nil {
		return nil, fmt.Errorf("identity: %w", err)
	}
	if err := user.Create(ip); err != nil {
		return nil, fmt.Errorf("identity: %w", err)
	}
	dev.User = *user

	return dev, nil
}

func MakeIdentity(ip string) (*models.Device, error) {
	dev := new(models.Device)
	dev.IP = ip

	if err := dev.Identify(); err != nil {
		return nil, err
	}

	return dev, nil
}
