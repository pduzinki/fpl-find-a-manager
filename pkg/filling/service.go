package filling

import (
	"fpl-find-a-manager/pkg/adding"
	"fpl-find-a-manager/pkg/listing"
)

//
type Service interface {
	Fill()
}

type service struct {
	adder  *adding.Service
	lister *listing.Service
}

//
func NewService(adder adding.Service, lister listing.Service) Service {
	return &service{}
}

//
func (s *service) Fill() {

}
