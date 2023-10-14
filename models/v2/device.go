package models

import (
	"fmt"
	"net"

	"go.mongodb.org/mongo-driver/bson"

	"application/data"
)

type Device struct {
	IP      string  `json:"ip"`
	Company Company `json:"company"`
	Branch  Branch  `json:"branch"`
	User    User    `json:"user"`
	Type    string  `json:"type"`
}

const devices = "devices"

func DeviceGetByIP(ip string) (*Device, error) {
	coll := data.GetCollection(devices)

	var dev Device
	if err := coll.FindOne(nil, bson.D{{Key: "ip", Value: ip}}).Decode(&dev); err != nil {
		return nil, err
	}

	return &dev, nil
}

func DevicesAll() ([]Device, error) {
	coll := data.GetCollection(devices)

	cursor, err := coll.Find(nil, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("all: %w", err)
	}

	var devices []Device
	if err := cursor.All(nil, &devices); err != nil {
		return nil, fmt.Errorf("all: %w", err)
	}

	return devices, nil
}

func (d *Device) Exists() (bool, error) {
	coll := data.GetCollection(devices)

	count, err := coll.CountDocuments(nil, bson.D{{Key: "ip", Value: d.IP}})
	if err != nil {
		return false, fmt.Errorf("exists: %w", err)
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (d *Device) Create() error {
	coll := data.GetCollection(devices)

	if d.IP == "" {
		return fmt.Errorf("create: ip is empty")
	}

	if exists, err := d.Exists(); err != nil {
		return err
	} else if exists {
		return fmt.Errorf("create: %s already exists", d.IP)
	}

	if _, err := coll.InsertOne(nil, d); err != nil {
		return fmt.Errorf("create: %w", err)
	}

	return nil
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

func (d *Device) Identify() error {
	db := data.DB

	var networks []Network
	if err := db.Select(&networks, "SELECT * FROM networks ORDER by netmask"); err != nil {
		return err
	}
	var network Network
	for _, n := range networks {
		_, ipnet, err := net.ParseCIDR(n.Netmask)
		size, _ := ipnet.Mask.Size()
		ip, _, err := net.ParseCIDR(d.IP + "/" + fmt.Sprint(size))
		if err != nil {
			return err
		}
		if ipnet.Contains(ip) {
			network = n
		}
	}

	var branch Branch
	if err := db.Get(&branch, "SELECT * FROM branches WHERE id = $1", network.BranchID); err != nil {
		return err
	}

	d.Branch = branch

	var company Company
	if err := db.Get(&company, "SELECT * FROM companies WHERE id = $1", branch.CompanyID); err != nil {
		return err
	}

	d.Company = company

	return nil
}
