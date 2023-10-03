package models

import "errors"

var (
	ErrUserNotFound      = errors.New("system: user not found for device")
	ErrUnsupportedDevice = errors.New("system: device's ip not found in subnets")
)
