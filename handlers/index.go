package handlers

import (
	"application/data"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleTest(c *fiber.Ctx) error {
    issues, _ := data.DB.ListIssues()
    return c.Render("test", fiber.Map{
        "Issues": issues,
    })
}

func HandleTestChangeStatus(c *fiber.Ctx) error {
    clientId := c.Params("client_id")
    issueId := c.Params("issue_id")

    action := c.Query("action")

    issue, _ := data.DB.RetrieveIssue(issueId, clientId)
    if action == "accept" {
        issue.Status = "Принята к исполнению"
    }
    if action == "decline" {
        issue.Status = "Отклонена"
    }
    res, err := data.DB.UpdateIssue(issue)
    
    log.Printf("Issue: %v\nRows: %v\nError: %v", issue, res, err)

    return nil
}

func HandleIndex(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{
        "ClientID": GetClientID(c),
    })
}

func HandleIssues(c *fiber.Ctx) error {
    clientID := GetClientID(c)
    issues, _ := data.DB.ListClientIssues(clientID)
    return c.Render("issue", fiber.Map{
        "Issues": issues,
    })
}


func HandleIssueSend(c *fiber.Ctx) error {
    clientID := GetClientID(c)
    issue := data.NewIssue(clientID)
    issue.Status = "Новое обращение"
    if err := c.BodyParser(issue); err != nil {
        return err
    }

    data.DB.InsertIssue(*issue)
    log.Printf("formData: %v", issue)
    return c.SendStatus(200)
}
