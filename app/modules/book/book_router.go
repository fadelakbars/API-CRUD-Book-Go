package book

import (
	"book-management/app/modules/book/controller"
	"book-management/app/modules/book/repository"
	"book-management/app/modules/book/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(app fiber.Router, db *gorm.DB) {
	repository := repository.NewBookRepository()
	service := service.NewBookService(repository, db)
	controller := controller.NewBookController(service)

	route := app.Group("/books")

	route.Get("/", controller.GetAllBook) 
	route.Get("/:id", controller.GetBookByID) 
	route.Post("/", controller.CreateBook)
	route.Put("/:id", controller.UpdateBook)
	route.Delete("/:id", controller.DeleteBook)

}