package sqlite

import "gorm.io/gorm"

//
type Manager struct {
	gorm.Model
	FplID    int
	FullName string
}
