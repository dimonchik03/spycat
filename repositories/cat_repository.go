package repositories

import (
	"errors"
	"gorm.io/gorm"
	"spycat/models"
)

type CatRepository interface {
	Create(cat *models.Cat) error
	FindAll() ([]models.Cat, error)
	FindByID(id uint) (*models.Cat, error)
	Update(cat *models.Cat) error
	Delete(id uint) error
}

type catRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) CatRepository {
	return &catRepository{db: db}
}

func (r *catRepository) Create(cat *models.Cat) error {
	return r.db.Create(cat).Error
}

func (r *catRepository) FindAll() ([]models.Cat, error) {
	var cats []models.Cat
	err := r.db.Find(&cats).Error
	return cats, err
}

func (r *catRepository) FindByID(id uint) (*models.Cat, error) {
	var cat models.Cat
	err := r.db.First(&cat, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("cat not found")
		}
		return nil, err
	}
	return &cat, nil
}

func (r *catRepository) Update(cat *models.Cat) error {
	return r.db.Save(cat).Error
}

func (r *catRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Cat{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("cat not found")
	}
	return nil
}
