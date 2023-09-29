package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func HandleChat(c *fiber.Ctx) error {
	return c.Render("pages/chat", fiber.Map{})
}

func HandleWebSocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func HandleGetWebSocket(c *websocket.Conn) {
	// mt  int
	var msg []byte
	var err error
	// id := c.Params("id")
	for {
		if _, msg, err = c.ReadMessage(); err != nil {
			log.Println("read: ", err)
			break
		}
		log.Printf("recv: %s", msg)
		html := fmt.Sprintf(`<!-- will be interpreted as hx-swap-oob="true" by default -->
<form id="form">
    ...
</form>
<!-- will be appended to #notifications div -->
<div id="notifications" hx-swap-oob="beforeend">
    New message received
</div>
<!-- will be swapped using an extension -->
<div id="chat_room" hx-swap-oob="morphdom">
    ....
</div>%s`, msg)

		if err = c.WriteMessage(websocket.BinaryMessage, []byte(html)); err != nil {
			log.Println("write: ", err)
			break
		}
		time.Sleep(time.Second * 2)
	}
}
