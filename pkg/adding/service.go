package adding

import "fmt"

//
type Service interface {
	AddManager(manager Manager)
}

//
type Repository interface {
	AddManager(manager Manager) error
}

type service struct {
	r Repository
}

//
func NewService(r Repository) Service {
	return &service{
		r: r,
	}
}

//
func (s *service) AddManager(manager Manager) {
	err := s.r.AddManager(manager)
	if err != nil {
		fmt.Println(err)
	}
}
