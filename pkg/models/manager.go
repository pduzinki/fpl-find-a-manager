package models

import "gorm.io/gorm"

//
type Manager struct {
	gorm.Model
	FplID              int
	FullName           string
	FullNameNormalized string
}
