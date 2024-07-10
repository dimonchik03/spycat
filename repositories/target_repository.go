package repositories

import (
	"gorm.io/gorm"
	"spycat/spycat/models"
)

type TargetRepository struct {
	DB *gorm.DB
}

func NewTargetRepository(db *gorm.DB) TargetRepository {
	return TargetRepository{DB: db}
}

func (r *TargetRepository) Create(target *models.Target) error {
	return r.DB.Create(target).Error
}

func (r *TargetRepository) Update(target *models.Target) error {
	return r.DB.Save(target).Error
}

func (r *TargetRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Target{}, id).Error
}

func (r *TargetRepository) FindByID(id uint) (*models.Target, error) {
	var target models.Target
	err := r.DB.Preload("Mission.Cat").First(&target, id).Error
	return &target, err
}

func (r *TargetRepository) FindByMissionID(missionID uint) ([]models.Target, error) {
	var targets []models.Target
	err := r.DB.Where("mission_id = ?", missionID).Preload("Mission.Cat").Find(&targets).Error
	return targets, err
}

func (r *TargetRepository) FindAll() ([]models.Target, error) {
	var targets []models.Target
	err := r.DB.Preload("Mission.Cat").Find(&targets).Error
	return targets, err
}
