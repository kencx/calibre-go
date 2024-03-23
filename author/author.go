package author

import (
	"context"

	"github.com/kencx/calibre-go/internal/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func GetById(id int) (*models.Author, error) {
	return models.Authors(qm.Where("id = ?", id)).One(context.Background(), boil.GetContextDB())
}

func GetByName(name string) (*models.Author, error) {
	return models.Authors(qm.Where("name LIKE ?", name)).One(context.Background(), boil.GetContextDB())
}
