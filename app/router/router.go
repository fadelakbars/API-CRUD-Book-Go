package router

import (
	"book-management/app/modules/book"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) { 
	apiRoutes := app.Group("/api")

	book.Router(apiRoutes, db) 
}