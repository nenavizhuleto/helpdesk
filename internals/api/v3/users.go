package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"helpdesk/internals/megaplan"
	"helpdesk/internals/models/v2"
	"helpdesk/internals/models/v3/task"
	"helpdesk/internals/models/v3/user"
)

func SetUsersRoutes(path string, router fiber.Router) {
	users := router.Group(path)
	users.Get("/", GetUsers)
	users.Get("/:id", GetUser)
	users.Post("/", CreateUser)
	users.Put("/", UpdateUser)
	users.Delete("/", DeleteUser)

	users.Get("/:id/tasks", GetUserTasks)
	users.Post("/:id/tasks", CreateUserTask)
}

func GetUsers(c *fiber.Ctx) error {
	users, err := user.All()
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := user.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	
	var body struct {
		Name string
		Phone string
	}
	if err := c.BodyParser(&body); err != nil {
		return fmt.Errorf("createUser: %w", err)
	}
	user, err := user.New(body.Name, body.Phone)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return fmt.Errorf("updateUser: %w", err)
	}

	if err := user.Update(); err != nil {
		return fmt.Errorf("updateUser: %w", err)
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := user.Get(id)
	if err != nil {
		return err
	}

	if err := user.Delete(); err != nil {
		return err
	}

	return c.JSON(user)
}

func GetUserTasks(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := user.Get(id)
	if err != nil {
		return err
	}

	tasks, err := task.ByUser(user)
	if err != nil {
		return err
	}

	return c.JSON(tasks)

}

func CreateUserTask(c *fiber.Ctx) error {
	id := c.Params("id")

	dev, err := models.DeviceGetByIP(c.IP())
	if err != nil {
		return fmt.Errorf("user: %w", err)
	}

	var task = models.NewTask()

	if err := c.BodyParser(&task); err != nil {
		return fmt.Errorf("user: %w", err)
	}

	user, err := models.UserByID(id)
	if err != nil {
		return err
	}

	task.User = *user
	task.Company = dev.Company
	task.Branch = dev.Branch

	task.BeforeCreateHook = func(t *models.Task) error {
		var TaskSubjectFormat = `
			<h2>от %s:</h2>
			<h3>Суть обращения:</h3>
			<p>%s</p>
			<hr/>
			<h3>Дополнительная информания:</h3>
			<ul>
			<li>Контакты: %s</li>
			<li>Устройство: %s</li>
			<li>Отдел: <br/>Название: %s <br/>Описание: %s <br/>Адрес: %s <br/>Контакты: %s</li>
			</ul>
		`
		task_name := fmt.Sprintf("%s: %s", t.Company.Name, t.Name)
		task_subject := fmt.Sprintf(TaskSubjectFormat,
			t.User.Name,
			t.Subject,
			t.User.Phone,
			t.User.Devices[0],
			t.Branch.Name,
			t.Branch.Description,
			t.Branch.Address,
			t.Branch.Contacts,
		)

		dto, err := megaplan.MP.CreateTask(task_name, task_subject)
		if err != nil {
			return fmt.Errorf("before_create_hook: %w", err)
		}

		if dto.TimeCreated != nil {
			t.ID = dto.ID
			t.TimeCreated = dto.TimeCreated.Value
			t.LastActivity = dto.TimeCreated.Value
		}

		return nil
	}

	if err := task.Create(); err != nil {
		return err
	}

	return c.JSON(task)

}
