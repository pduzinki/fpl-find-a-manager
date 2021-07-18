package models

import (
	"fmt"
	"fpl-find-a-manager/pkg/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//
type ManagerDB interface {
	AddManager(manager *Manager) error
	AddManagers(managers []Manager) error
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
func NewManagerService(cfg config.DatabaseConfig) (ManagerService, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, 5432, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	db.Migrator().DropTable(&Manager{}) // TODO remove later
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
func (ms *managerService) AddManagers(managers []Manager) error {
	return ms.ManagerDB.AddManagers(managers)
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

func (mv *managerValidator) AddManagers(managers []Manager) error {
	// TODO add validations
	return mv.ManagerDB.AddManagers(managers)
}

func (mv *managerValidator) MatchManagersByName(name string) ([]Manager, error) {
	m := Manager{}
	m.FullName = name
	err := runManagerValidatorFuncs(&m, mv.FullNameLongerThanThreeRunes)
	if err != nil {
		return nil, err
	}

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

func (mg *managerGorm) AddManagers(managers []Manager) error {
	return mg.db.Create(managers).Error
}

func (mg *managerGorm) MatchManagersByName(name string) ([]Manager, error) {
	managers := make([]Manager, 0)

	err := mg.db.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Find(&managers).Error

	return managers, err
}
