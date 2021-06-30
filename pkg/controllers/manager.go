package controllers

import (
	"fmt"
	"fpl-find-a-manager/pkg/models"
	"fpl-find-a-manager/pkg/wrapper"
)

//
type ManagerController struct {
	ms models.ManagerService
	w  wrapper.Wrapper
}

//
func NewManagerController(ms models.ManagerService) *ManagerController {
	return &ManagerController{
		ms: ms,
		w:  wrapper.NewWrapper(),
	}
}

// MatchManagersByName returns managers that name contains given string.
func (mc *ManagerController) MatchManagersByName(name string) ([]models.Manager, error) {
	managers, err := mc.ms.MatchManagersByName(name)
	if err != nil {
		return nil, err
	}

	return managers, nil
}

// AddManagers constantly queries FPL API wrapper and keeps managers db up-to-date.
// Should be run in goroutine.
func (mc *ManagerController) AddManagers() {
	wm, err := mc.w.GetManager(1239)
	if err != nil {
		fmt.Println("failed to get manager via fpl api")
	}

	am := models.Manager{
		FplID:    wm.ID,
		FullName: fmt.Sprintf("%s %s", wm.FirstName, wm.LastName),
	}

	mc.ms.AddManager(&am)

	fmt.Println(wm)
}
