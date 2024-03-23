package book

import (
	"context"

	"github.com/kencx/calibre-go/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetById(id int) (*models.Book, error) {
	return models.Books(qm.Where("id = ?", id)).One(context.Background(), boil.GetContextDB())
}

func GetByTitle(title string) (*models.Book, error) {
	return models.Books(qm.Where("title LIKE ?", title)).One(context.Background(), boil.GetContextDB())
}

func GetByAuthor(authorName string) (models.BookSlice, error) {
	return models.Books(
		qm.Select("*"),
		qm.InnerJoin("books_authors_link bal on bal.book = books.id"),
		qm.InnerJoin("authors ON authors.id = bal.author"),
		qm.Where("authors.name LIKE ?", authorName),
	).All(context.Background(), boil.GetContextDB())
}
