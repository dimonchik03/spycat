package repositories

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"spycat/spycat/models"
)

type MissionRepository interface {
	Create(mission *models.Mission) error
	FindAll() ([]models.Mission, error)
	FindByID(id uint) (*models.Mission, error)
	Update(mission *models.Mission) error
	Delete(id uint) error
}

type missionRepository struct {
	db *gorm.DB
}

func NewMissionRepository(db *gorm.DB) MissionRepository {
	return &missionRepository{db: db}
}

func (r *missionRepository) Create(mission *models.Mission) error {
	return r.db.Create(mission).Error
}

func (r *missionRepository) FindAll() ([]models.Mission, error) {
	var missions []models.Mission
	err := r.db.Preload("Targets").Preload("Cat").Find(&missions).Error
	if err != nil {
		log.Printf("error fetching missions: %v", err)
	}

	return missions, err
}

func (r *missionRepository) FindByID(id uint) (*models.Mission, error) {
	var mission models.Mission
	err := r.db.Preload("Targets").Preload("Cat").First(&mission, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("mission not found")
		}
		return nil, err
	}
	return &mission, nil
}

func (r *missionRepository) Update(mission *models.Mission) error {
	return r.db.Save(mission).Error
}

func (r *missionRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Mission{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("mission not found")
	}
	return nil
}
