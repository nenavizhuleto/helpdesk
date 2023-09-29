package auth

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"

	"application/models"
)

type Identity struct {
	User struct {
		ID    string
		Name  string
		Phone string
	}
	Device struct {
		IP   string
		Type string
	}
	Subnet struct {
		Network string
	}
	Branch struct {
		Name        string
		Address     string
		Contacts    string
		Description string
	}
	Company struct {
		ID   string
		Name string
		Slug string
	}
}

var ErrUserNotFound = errors.New("auth: user not found for device")

func GetIdentity(c *fiber.Ctx) *Identity {
	i := c.Locals("Identity").(*Identity)
	return i
}

func MakeIdentity(ip string) (*Identity, error) {
	i := &Identity{}

	log.Println("ID: device")
	d, err := models.NewDevice(ip)
	if err != nil {
		return nil, err
	}
	log.Println(d)

	i.Device.IP = d.IP
	i.Device.Type = d.Type

	log.Println("ID: user")
	u, err := models.GetUserFromDevice(d)
	if err != nil {
		return nil, ErrUserNotFound
	}
	log.Println(u)

	i.User.ID = u.ID
	i.User.Name = u.Name
	i.User.Phone = u.Phone

	log.Println("ID: subnet")
	s, err := models.GetSubnetFromDevice(d)
	if err != nil {
		return nil, err
	}

	i.Subnet.Network = s.Netmask

	log.Println("ID: branch")
	b, err := models.GetBranchFromSubnet(s)
	if err != nil {
		return nil, err
	}
	log.Println(b)

	i.Branch.Name = b.Name
	i.Branch.Address = b.Address
	i.Branch.Contacts = b.Contacts
	i.Branch.Description = b.Description

	log.Println("ID: company")
	c, err := models.GetCompanyFromBranch(b)
	if err != nil {
		return nil, err
	}
	log.Println(c)

	i.Company.ID = c.ID
	i.Company.Name = c.Name
	i.Company.Slug = c.Slug

	log.Println(i)

	return i, nil
}
