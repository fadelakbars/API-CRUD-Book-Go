package repository

import (
	"book-management/app/domain"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (r *BookRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.Book, error) {
	var books []domain.Book
	err := db.WithContext(ctx).Find(&books).Error
	return books, err
}

func (r *BookRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Book, error) {
	var book domain.Book
	err := db.WithContext(ctx).First(&book, "id = ?", id).Error
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

func (r *BookRepositoryImpl) Create(ctx context.Context, db *gorm.DB, book domain.Book) (domain.Book, error) {
	err := db.WithContext(ctx).Create(&book).Error
	return book, err
}

func (r *BookRepositoryImpl) Update(ctx context.Context, db *gorm.DB, book domain.Book) (domain.Book, error) {
	err := db.WithContext(ctx).Save(&book).Error
	return book, err
}

func (r *BookRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	return db.WithContext(ctx).Delete(&domain.Book{}, "id = ?", id).Error
}