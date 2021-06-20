package sqlite

import (
	"fmt"
	"fpl-find-a-manager/pkg/adding"
	"fpl-find-a-manager/pkg/listing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//
type Storage struct {
	db *gorm.DB
}

//
func NewStorage() (*Storage, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Manager{})

	return &Storage{
		db: db,
	}, nil
}

//
func (s *Storage) AddManager(manager adding.Manager) error {
	sm := Manager{
		FplID:    manager.FplID,
		FullName: manager.FullName,
	}

	return s.db.Create(&sm).Error
}

//
func (s *Storage) GetManagerByName(name string) (*listing.Manager, error) {
	manager := listing.Manager{}

	err := s.db.Where("full_name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		First(&manager).Error
	// TODO remove First(), and return a slice instead

	return &manager, err
}
