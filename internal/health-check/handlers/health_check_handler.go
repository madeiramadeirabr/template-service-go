package healthCheck

import "github.com/gofiber/fiber/v2"

func GetStatus(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]interface{}{
		"status": "OK",
	})
}
