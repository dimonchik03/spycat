package services

import (
	"spycat/spycat/models"
	"spycat/spycat/repositories"
)

type CatService struct {
	repo repositories.CatRepository
}

func NewCatService(repo repositories.CatRepository) *CatService {
	return &CatService{repo: repo}
}

func (s *CatService) CreateCat(cat *models.Cat) error {
	if err := validateBreed(cat.Breed); err != nil {
		return err
	}

	return s.repo.Create(cat)
}

func (s *CatService) GetCats() ([]models.Cat, error) {
	return s.repo.FindAll()
}

func (s *CatService) GetCat(id uint) (*models.Cat, error) {
	return s.repo.FindByID(id)
}

func (s *CatService) UpdateCat(id uint, updatedCat *models.Cat) error {
	cat, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	cat.Salary = updatedCat.Salary

	return s.repo.Update(cat)
}

func (s *CatService) DeleteCat(id uint) error {
	return s.repo.Delete(id)
}

func validateBreed(breed string) error {
	return nil
}
