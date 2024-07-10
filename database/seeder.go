// File: database/seeder.go

package database

import (
	models2 "spycat/spycat/models"

	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {

	var userCount int64
	db.Model(&models2.User{}).Count(&userCount)
	if userCount == 0 {
		users := []models2.User{
			{Username: "admin", Password: "password"},
		}
		if err := db.Create(&users).Error; err != nil {
			return err
		}
	}

	// Seed Cats
	var catCount int64
	db.Model(&models2.Cat{}).Count(&catCount)
	if catCount == 0 {
		cats := []models2.Cat{
			{Name: "Whiskers", YearsOfExperience: 3, Breed: "Siamese", Salary: 50000},
			{Name: "Shadow", YearsOfExperience: 5, Breed: "Persian", Salary: 60000},
			{Name: "Mittens", YearsOfExperience: 2, Breed: "Maine Coon", Salary: 45000},
			{Name: "Tiger", YearsOfExperience: 4, Breed: "Bengal", Salary: 55000},
			{Name: "Luna", YearsOfExperience: 1, Breed: "Sphynx", Salary: 40000},
		}
		if err := db.Create(&cats).Error; err != nil {
			return err
		}
	}

	// Seed Missions and Targets
	var missionCount int64
	db.Model(&models2.Mission{}).Count(&missionCount)
	if missionCount == 0 {
		// Create missions and associated targets
		missions := []models2.Mission{
			{CatID: 1, Complete: false, Targets: []models2.Target{
				{Name: "Target1", Country: "CountryA", Notes: "Initial notes", Complete: false},
				{Name: "Target2", Country: "CountryB", Notes: "Initial notes", Complete: false},
			}},
			{CatID: 2, Complete: false, Targets: []models2.Target{
				{Name: "Target3", Country: "CountryC", Notes: "Initial notes", Complete: false},
				{Name: "Target4", Country: "CountryD", Notes: "Initial notes", Complete: false},
			}},
			{CatID: 3, Complete: false, Targets: []models2.Target{
				{Name: "Target5", Country: "CountryE", Notes: "Initial notes", Complete: false},
				{Name: "Target6", Country: "CountryF", Notes: "Initial notes", Complete: false},
			}},
			{CatID: 4, Complete: false, Targets: []models2.Target{
				{Name: "Target7", Country: "CountryG", Notes: "Initial notes", Complete: false},
				{Name: "Target8", Country: "CountryH", Notes: "Initial notes", Complete: false},
			}},
			{CatID: 5, Complete: false, Targets: []models2.Target{
				{Name: "Target9", Country: "CountryI", Notes: "Initial notes", Complete: false},
				{Name: "Target10", Country: "CountryJ", Notes: "Initial notes", Complete: false},
			}},
		}
		for _, mission := range missions {
			if err := db.Create(&mission).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
