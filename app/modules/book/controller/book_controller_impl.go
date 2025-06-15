package controller

import (
	"book-management/app/domain"
	"book-management/app/helper"
	"book-management/app/modules/book/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BookControllerImpl struct {
	service service.BookService
}
func NewBookController(service service.BookService) BookController {
	return &BookControllerImpl{service: service}
}

func (c *BookControllerImpl) GetAllBook(ctx *fiber.Ctx) error {
	books, err := c.service.FindAll(ctx.Context())
	if err != nil {
		return helper.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to fetch books")
	}
	return helper.WriteJson(ctx, fiber.StatusOK, "Books fetched successfully", books)
}

func (c *BookControllerImpl) GetBookByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))

	if err != nil {
		return helper.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid book ID")
	}
	book, err := c.service.FindByID(ctx.Context(), id)
	if err != nil {
		return helper.HandleError(ctx, err, fiber.StatusNotFound, "Book not found")
	}
	return helper.WriteJson(ctx, fiber.StatusOK, "Book fetched successfully", book)
}

func (c *BookControllerImpl) CreateBook(ctx *fiber.Ctx) error {
	var book domain.Book
	if err := ctx.BodyParser(&book); err != nil {
		return helper.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request body")
	}
	book, err := c.service.Create(ctx.Context(), book)
	if err != nil {
		return helper.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to create book")
	}
	return helper.WriteJson(ctx, fiber.StatusCreated, "Book created successfully", book)
}

func (c *BookControllerImpl) UpdateBook(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helper.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid book ID")
	}

	var book domain.Book
	if err := ctx.BodyParser(&book); err != nil {
		return helper.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid request body")
	}
	book.ID = id 
	updatedBook, err := c.service.Update(ctx.Context(), book)
	if err != nil {
		return helper.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to update book")
	}
	return helper.WriteJson(ctx, fiber.StatusOK, "Book updated successfully", updatedBook)
}

func (c *BookControllerImpl) DeleteBook(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return helper.HandleError(ctx, err, fiber.StatusBadRequest, "Invalid book ID")
	}

	if err := c.service.Delete(ctx.Context(), id); err != nil {
		return helper.HandleError(ctx, err, fiber.StatusInternalServerError, "Failed to delete book")
	}
	return helper.WriteJson(ctx, fiber.StatusOK, "Book deleted successfully", nil)
}