package handlers

import (
	"application/data"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleTest(c *fiber.Ctx) error {
    return c.Render("test", fiber.Map{
        "Map": data.GetDb(),
    })
}

func HandleTestChangeStatus(c *fiber.Ctx) error {
    id := c.Params("uuid")
    index := c.Params("index")

    issues, _ := data.GetIssuesById(id)
    issue := issues[index]
    log.Printf("Index: %s\n\nIssues: %v\n\nIssue: %v\n\n", index, issues, issue)
    issue.Status = "Принята к исполнению"
    data.StoreIssue(id, &issue)

    return nil
}

func HandleIndex(c *fiber.Ctx) error {
    id := c.Cookies("uuid")
    issues, _ := data.GetIssuesById(id)
    return c.Render("index", fiber.Map{
        "Uuid": id,
        "Issues": issues,
    })
}

func HandleIssues(c *fiber.Ctx) error {
    id := c.Cookies("uuid") 
    issues := data.D.Get(id)
    log.Printf("Event: %v", issues)
    return c.Render("issue", fiber.Map{
        "Issues": issues,
    })
}


func HandleIssueSend(c *fiber.Ctx) error {
    id := c.Cookies("uuid")
    formData := &data.Issue{
        ID: uuid.NewString(),
        Status: "Новое обращение",
    }
    if err := c.BodyParser(formData); err != nil {
        return err
    }

    data.StoreIssue(id, formData)
    log.Printf("formData: %v", formData)
    return c.SendStatus(200)
}
