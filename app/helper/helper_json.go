package helper

import "github.com/gofiber/fiber/v2"

type Response struct{
	Status int			`json:"status"`
	Message string		`json:"message"`
	Data   interface{}	`json:"data,omitempty"`
}

func WriteJson(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}