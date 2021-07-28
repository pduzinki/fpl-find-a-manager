package models

import "gorm.io/gorm"

//
type Manager struct {
	gorm.Model
	FplID              int
	FullName           string
	FullNameNormalized string
}

type Managers []Manager

func (m Managers) Len() int {
	return len(m)
}

func (m Managers) Less(i, j int) bool {
	return (m[i].FplID < m[j].FplID)
}

func (m Managers) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
