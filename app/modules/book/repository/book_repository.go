package repository

import (
	"book-management/app/domain"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.Book, error)
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Book, error)
	Create(ctx context.Context, db *gorm.DB, book domain.Book) (domain.Book, error)
	Update(ctx context.Context, db *gorm.DB, book domain.Book) (domain.Book, error)
	Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error
}