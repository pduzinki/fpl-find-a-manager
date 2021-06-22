package adding

import (
	"fmt"
	"time"

	"fpl-find-a-manager/pkg/listing"
	"fpl-find-a-manager/pkg/wrapper"
)

//
type Service interface {
	AddManager(manager Manager) // TODO remove it later
	AddAllManagers()
}

//
type Repository interface {
	AddManager(manager Manager) error
	GetLastManager() (*listing.Manager, error) // TODO clear this mess
}

type service struct {
	w      wrapper.Wrapper
	r      Repository
	lastID int
}

//
func NewService(r Repository) Service {
	return &service{
		w:      wrapper.NewWrapper(),
		r:      r,
		lastID: 0,
	}
}

//
func (s *service) AddManager(manager Manager) {
	err := s.r.AddManager(manager)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *service) AddAllManagers() {
	manager, err := s.r.GetLastManager()
	if err != nil {
		s.lastID = 0
	} else {
		s.lastID = manager.FplID
	}

	fmt.Printf("Last ID is: %d\n", s.lastID)

	start := time.Now()
	for id := s.lastID + 1; id < 10_000_000; id++ {
		wm, err := s.w.GetManager(id)
		if err != nil {
			return
		}

		am := Manager{
			FplID:    wm.ID,
			FullName: fmt.Sprintf("%s %s", wm.FirstName, wm.LastName),
		}

		s.AddManager(am)
	}
	duration := time.Since(start)
	fmt.Println(duration)
}
