package controllers

import "fpl-find-a-manager/pkg/models"

//
type ManagerController struct {
}

//
func NewManagerController(ms models.ManagerService) *ManagerController {
	return &ManagerController{}
}

// MatchManagersByName returns managers that name contains given string.
func (mc *ManagerController) MatchManagersByName(name string) ([]models.Manager, error) {
	return nil, nil
}

// AddManagers constantly queries FPL API wrapper and keeps managers db up-to-date.
// Should be run in goroutine.
func (mc *ManagerController) AddManagers() {

}
