package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ManagerService interface {
	AddManager(manager *Manager) error
	MatchManagersByName(name string) ([]Manager, error)
}

type managerService struct {
	db *gorm.DB
}

func NewManagerService() (ManagerService, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Manager{})

	return &managerService{
		db: db,
	}, nil
}

func (ms *managerService) AddManager(manager *Manager) error {
	return ms.db.Create(&manager).Error
}

func (ms *managerService) MatchManagersByName(name string) ([]Manager, error) {
	managers := make([]Manager, 0)

	err := ms.db.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Find(&managers).Error

	return managers, err
}
