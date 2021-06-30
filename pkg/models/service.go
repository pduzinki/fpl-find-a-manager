package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//
type ManagerDB interface {
	AddManager(manager *Manager) error
	MatchManagersByName(name string) ([]Manager, error)
}

//
type ManagerService interface {
	ManagerDB
}

type managerService struct {
	ManagerDB
}

//
func NewManagerService() (ManagerService, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Manager{})

	mg := newManagerGorm(db)
	mv := newManagerValidator(mg)

	return &managerService{
		ManagerDB: mv,
	}, nil
}

//
func (ms *managerService) AddManager(manager *Manager) error {
	return ms.ManagerDB.AddManager(manager)
}

//
func (ms *managerService) MatchManagersByName(name string) ([]Manager, error) {
	return ms.ManagerDB.MatchManagersByName(name)
}

// ----------------------------------------------------------------
type managerValidator struct {
	ManagerDB
}

func newManagerValidator(m ManagerDB) *managerValidator {
	return &managerValidator{
		ManagerDB: m,
	}
}

func (mv *managerValidator) AddManager(manager *Manager) error {
	// TODO add validations
	return mv.ManagerDB.AddManager(manager)
}

func (mv *managerValidator) MatchManagersByName(name string) ([]Manager, error) {
	// TODO add validations
	return mv.ManagerDB.MatchManagersByName(name)
}

// ----------------------------------------------------------------
type managerGorm struct {
	db *gorm.DB
}

func newManagerGorm(db *gorm.DB) *managerGorm {
	return &managerGorm{
		db: db,
	}
}

func (mg *managerGorm) AddManager(manager *Manager) error {
	return mg.db.Create(&manager).Error
}

func (mg *managerGorm) MatchManagersByName(name string) ([]Manager, error) {
	managers := make([]Manager, 0)

	err := mg.db.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Find(&managers).Error

	return managers, err
}
