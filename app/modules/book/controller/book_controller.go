package controller

import "github.com/gofiber/fiber/v2"

type BookController interface {
	GetAllBook(c *fiber.Ctx) error
	GetBookByID(c *fiber.Ctx) error
	CreateBook(c *fiber.Ctx) error
	UpdateBook(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
}