// File: services/target_service.go

package services

import (
	"errors"
	"spycat/spycat/models"
	repositories2 "spycat/spycat/repositories"
)

type TargetService struct {
	targetRepo  repositories2.TargetRepository
	missionRepo repositories2.MissionRepository
}

func NewTargetService(targetRepo repositories2.TargetRepository, missionRepo repositories2.MissionRepository) *TargetService {
	return &TargetService{
		targetRepo:  targetRepo,
		missionRepo: missionRepo,
	}
}

func (s *TargetService) AddTarget(target *models.Target) error {
	mission, err := s.missionRepo.FindByID(target.MissionID)
	if err != nil {
		return err
	}

	if mission.Complete {
		return errors.New("cannot add target to completed mission")
	}

	return s.targetRepo.Create(target)
}

func (s *TargetService) UpdateTarget(target *models.Target) error {
	existingTarget, err := s.targetRepo.FindByID(target.ID)
	if err != nil {
		return err
	}

	if existingTarget.Complete {
		return errors.New("cannot update completed target")
	}

	return s.targetRepo.Update(target)
}

func (s *TargetService) DeleteTarget(id uint) error {
	target, err := s.targetRepo.FindByID(id)
	if err != nil {
		return err
	}

	if target.Complete {
		return errors.New("cannot delete completed target")
	}

	return s.targetRepo.Delete(id)
}

func (s *TargetService) CompleteTarget(id uint) error {
	target, err := s.targetRepo.FindByID(id)
	if err != nil {
		return err
	}

	target.Complete = true
	return s.targetRepo.Update(target)
}

func (s *TargetService) UpdateNotes(id uint, notes string) error {
	target, err := s.targetRepo.FindByID(id)
	if err != nil {
		return err
	}

	if target.Complete {
		return errors.New("cannot update notes for completed target")
	}

	mission, err := s.missionRepo.FindByID(target.MissionID)
	if err != nil {
		return err
	}

	if mission.Complete {
		return errors.New("cannot update notes for target in completed mission")
	}

	target.Notes = notes
	return s.targetRepo.Update(target)
}

func (s *TargetService) GetTargetsByMissionID(missionID uint) ([]models.Target, error) {
	return s.targetRepo.FindByMissionID(missionID)
}

func (s *TargetService) GetTargetByID(id uint) (*models.Target, error) {
	return s.targetRepo.FindByID(id)
}

func (s *TargetService) GetAllTargets() ([]models.Target, error) {
	return s.targetRepo.FindAll()
}
