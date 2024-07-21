// File: services/mission_service.go

package services

import (
	"errors"
	"spycat/models"
	"spycat/repositories"
)

type MissionService struct {
	repo repositories.MissionRepository
}

func NewMissionService(repo repositories.MissionRepository) *MissionService {
	return &MissionService{repo: repo}
}

func (s *MissionService) CreateMission(mission *models.Mission) error {
	if len(mission.Targets) < 1 || len(mission.Targets) > 3 {
		return errors.New("mission must have 1-3 targets")
	}
	return s.repo.Create(mission)
}

func (s *MissionService) GetMissions() ([]models.Mission, error) {
	return s.repo.FindAll()
}

func (s *MissionService) GetMission(id uint) (*models.Mission, error) {
	return s.repo.FindByID(id)
}

func (s *MissionService) UpdateMission(id uint, updatedMission *models.Mission) error {
	mission, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Update fields
	mission.CatID = updatedMission.CatID
	mission.Complete = updatedMission.Complete

	return s.repo.Update(mission)
}

func (s *MissionService) DeleteMission(id uint) error {
	mission, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if !mission.Complete {
		return s.repo.Delete(id)
	}
	return errors.New("cannot delete a completed mission")
}

func (s *MissionService) CompleteMission(id uint) error {
	mission, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	mission.Complete = true
	return s.repo.Update(mission)
}
