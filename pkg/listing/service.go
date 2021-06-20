package listing

import "fmt"

//
type Service interface {
	GetManagerByName(name string) (*Manager, error)
}

//
type Repository interface {
	GetManagerByName(name string) (*Manager, error)
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

func (s *service) GetManagerByName(name string) (*Manager, error) {
	m, err := s.r.GetManagerByName(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return m, nil
}
