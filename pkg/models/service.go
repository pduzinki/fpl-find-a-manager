package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ManagerService interface {
	Close()
}

type managerService struct {
}

func NewManagerService() ManagerService {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil
	}

	_ = db

	return nil
}

func (ms *managerService) Close() {

}
