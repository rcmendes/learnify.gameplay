package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/rcmendes/learnify-gameplay/adapters/repository/postgres/models"
	"github.com/rcmendes/learnify-gameplay/core/entities"
	"github.com/rcmendes/learnify-gameplay/core/ucs/ports"
)

type categoryRepository struct {
	connectFunc func() *pg.DB
}

//NewCategoryPostgresRepository creates a Category repository for Postgres database.
func NewCategoryPostgresRepository() ports.CategoryRepository {
	return &categoryRepository{
		connectFunc: Connect,
	}
}

//TODO Customize errors

func (repo *categoryRepository) Insert(category entities.Category) error {
	conn := repo.connectFunc()

	model := models.NewCategoryModel(category)

	if _, error := conn.Model(&model).Insert(); error != nil {
		return error
	}

	return nil
}

func (repo *categoryRepository) ListAll() (entities.CategoryList, error) {
	conn := repo.connectFunc()

	var modelList []*models.CategoryModel

	if err := conn.Model(&modelList).Select(); err != nil {
		return nil, err
	}

	list := make(entities.CategoryList, 0, len(modelList))

	for _, model := range modelList {
		c := model.ToEntity()
		list = append(list, c)
	}

	return list, nil
}

func (repo *categoryRepository) GetByName(name string) (*entities.Category, error) {
	conn := repo.connectFunc()

	var model models.CategoryModel

	if err := conn.Model(&model).
		Where("name = ?", name).
		Select(); err != nil {
		return nil, err
	}

	category := model.ToEntity()

	return category, nil
}
