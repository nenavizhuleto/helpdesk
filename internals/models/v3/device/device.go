package device

import (
	"go.mongodb.org/mongo-driver/bson"

	"helpdesk/internals/data"
	"helpdesk/internals/models"
	"helpdesk/internals/models/v3/branch"
	"helpdesk/internals/models/v3/company"
	"helpdesk/internals/models/v3/network"
	"helpdesk/internals/models/v3/user"
)

type Device struct {
	IP      string           `json:"ip"`
	Company *company.Company `json:"company"`
	Branch  *branch.Branch   `json:"branch"`
	Network *network.Network `json:"network"`
	User    *user.User       `json:"user"`
	Type    string           `json:"type"`
}

const devices = "devices"

var (
	PC      = "PC"
	Unknown = "unknown"
)

func New(ip string, company *company.Company, branch *branch.Branch, network *network.Network, user *user.User, typ string) (*Device, error) {
	validIP, err := newIP(ip)
	if err != nil {
		return nil, models.NewValidationError("device", "ip")
	}

	return &Device{
		IP:      validIP,
		Company: company,
		Branch:  branch,
		Network: network,
		User:    user,
		Type:    typ,
	}, nil
}

func Get(ip string) (*Device, error) {
	coll := data.GetCollection(devices)

	var dev Device
	if err := coll.FindOne(nil, bson.D{{Key: "ip", Value: ip}}).Decode(&dev); err != nil {
		return nil, models.NewDatabaseError("device", "get", err)
	}

	return &dev, nil
}

func All() ([]Device, error) {
	coll := data.GetCollection(devices)

	cursor, err := coll.Find(nil, bson.D{})
	if err != nil {
		return nil, models.NewDatabaseError("device", "all", err)
	}

	var devices []Device
	if err := cursor.All(nil, &devices); err != nil {
		return nil, models.NewDatabaseError("device", "all", err)
	}

	return devices, nil
}

// Updates all devices that a related to user
func DeviceUserCreateHook(u *user.User) error {
	coll := data.GetCollection(devices)
	for _, ip := range u.Devices {
		if err := coll.FindOneAndUpdate(nil, bson.D{{Key: "ip", Value: ip}}, bson.D{{Key: "$set", Value: bson.M{"user": u}}}).Err(); err != nil {
			return err
		}
	}

	return nil
}

func (d *Device) Save() error {
	coll := data.GetCollection(devices)

	if _, err := Get(d.IP); err != nil {
		// Not exists
		if _, err := coll.InsertOne(nil, d); err != nil {
			return models.NewDatabaseError("device", "create", err)
		}
	} else {
		// Exists
		if err := coll.FindOneAndUpdate(nil, bson.D{{Key: "ip", Value: d.IP}}, bson.D{{Key: "$set", Value: d}}).Err(); err != nil {
			return models.NewDatabaseError("device", "update", err)
		}

	}

	return nil
}

func Identify(ip string) error {

	network, err := network.GetByIP(ip)
	if err != nil {
		return err
	}

	branch, err := branch.Get(network.BranchID)
	if err != nil {
		return err
	}

	company, err := company.Get(branch.CompanyID)
	if err != nil {
		return err
	}

	device, err := New(ip, company, branch, network, nil, PC)
	if err != nil {
		return err
	}

	if err := device.Save(); err != nil {
		return err
	}

	return nil
}

// Private functions

func newIP(ip string) (string, error) {
	// Validate IP address
	return ip, nil
}
