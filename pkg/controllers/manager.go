package controllers

import (
	"fmt"
	"fpl-find-a-manager/pkg/models"
	"fpl-find-a-manager/pkg/wrapper"
	"time"
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
	// init worker pool
	// query fpl api via workers
	// collect 1000 entries (managers)
	// sort managers
	// put them in the database (in batch)
	// repeat until all managers are in the database

	totalManagers, err := mc.w.GetManagersCount()
	if err != nil {
		panic("shiet")
	}

	addedManagers := 0

	for totalManagers > addedManagers {
		start := time.Now()

		var numJobs = 1000
		if totalManagers-addedManagers < 1000 {
			numJobs = totalManagers - addedManagers
		}

		jobs := make(chan int, numJobs)
		results := make(chan models.Manager, numJobs)

		for w := 1; w <= 4; w++ {
			go mc.worker(w, jobs, results)
		}

		for j := 1 + addedManagers; j <= numJobs+addedManagers; j++ {
			// fmt.Println(j)
			jobs <- j
		}
		close(jobs)

		managers := make([]models.Manager, 0, 1_000)

		for a := 1; a <= numJobs; a++ {
			// tmp := <-results
			// fmt.Println(tmp)
			managers = append(managers, <-results)
		}

		duration := time.Since(start)
		fmt.Printf("It took %v to add 1000 fpl managers\n", duration)

		mc.ms.AddManagers(managers)

		addedManagers += numJobs
	}
}

func (mc *ManagerController) worker(id int, jobs <-chan int, results chan<- models.Manager) {
	for j := range jobs {
		wm, err := mc.w.GetManager(j)
		if err != nil {
			fmt.Println("failed to get manager via fpl api")
		}

		am := models.Manager{
			FplID:    wm.ID,
			FullName: fmt.Sprintf("%s %s", wm.FirstName, wm.LastName),
		}

		results <- am
	}
}
