package service

import (
	"book-management/app/domain"
	"book-management/app/modules/book/repository"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookService interface {
	FindAll(ctx context.Context) ([]domain.Book, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.Book, error)
	Create(ctx context.Context, book domain.Book) (domain.Book, error)
	Update(ctx context.Context, book domain.Book) (domain.Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type BookServiceImpl struct {
	repo repository.BookRepository
	db   *gorm.DB
}

func NewBookService(repo repository.BookRepository, db *gorm.DB) BookService {
	return &BookServiceImpl{repo: repo, db: db}
}

func (s *BookServiceImpl) FindAll(ctx context.Context) ([]domain.Book, error) {
	return s.repo.FindAll(ctx, s.db)
}

func (s *BookServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.Book, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *BookServiceImpl) Create(ctx context.Context, book domain.Book) (domain.Book, error) {
	return s.repo.Create(ctx, s.db, book)
}

func (s *BookServiceImpl) Update(ctx context.Context, book domain.Book) (domain.Book, error) {
	existingBook, err := s.repo.FindByID(ctx, s.db, book.ID)
	if err != nil {
		return book, err
	}

	book.CreatedAt = existingBook.CreatedAt
	return s.repo.Update(ctx, s.db, book)
}

func (s *BookServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, s.db, id)
}