package sqlite

//
type Manager struct {
	ID       int    `gorm:"not_null;index"`
	FullName string `gorm:"not_null;index"`
}
