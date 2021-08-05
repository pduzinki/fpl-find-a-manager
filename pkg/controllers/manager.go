package controllers

import (
	"fmt"
	"fpl-find-a-manager/pkg/models"
	"fpl-find-a-manager/pkg/wrapper"
	"log"
	"math/rand"
	"sort"
	"time"
)

var sleeps = []time.Duration{
	2 * time.Second,
	5 * time.Second,
	20 * time.Second,
	30 * time.Second,
}

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
	// collect N entries (managers)
	// sort N managers
	// put them in the database (in batch)
	// repeat until all managers are in the database

	totalManagers, err := mc.w.GetManagersCount()
	if err != nil {
		panic("shiet")
	}

	addedManagers := 0 // TODO use this later

	var numJobs = 1000
	jobs := make(chan int, numJobs)
	results := make(chan models.Manager, numJobs)

	for w := 1; w <= 64; w++ {
		go mc.worker(w, jobs, results)
	}

	for totalManagers > addedManagers {
		start := time.Now()

		if totalManagers-addedManagers < numJobs {
			numJobs = totalManagers - addedManagers
		}

		for j := 1 + addedManagers; j <= numJobs+addedManagers; j++ {
			jobs <- j
		}

		managers := make([]models.Manager, 0, numJobs)

		for len(managers) != numJobs {
			managers = append(managers, <-results)
		}

		duration := time.Since(start)
		fmt.Printf("################ It took %v to add %v fpl managers\n", duration, numJobs)

		sort.Sort(models.Managers(managers)) // so ID == fplID
		mc.ms.AddManagers(managers)

		addedManagers += numJobs
	}
}

func (mc *ManagerController) worker(id int, jobs chan int, results chan<- models.Manager) {
	w := wrapper.NewWrapper()

	for j := range jobs {
		wm, err := w.GetManager(j)
		if err == wrapper.ErrHTTPStatusNotFound {
			// no more managers to add, worker can be closed
			log.Println("FPL API call returned http 404, closing worker")
			return
		} else if err == wrapper.ErrHTTPTooManyRequests {
			// hit rate limit, let's sleep here for a bit, return job to pool
			jobs <- j
			sleepDuration := sleeps[rand.Intn(len(sleeps))]
			log.Println("FPL API call returned http 429, worker going to sleep for", sleepDuration)
			time.Sleep(sleepDuration)
			continue
		} else if err != nil {
			jobs <- j
			log.Printf("FPL API call returned %v, worker going to sleep.", err)
			time.Sleep(10 * time.Minute)
			continue
		}

		am := models.Manager{
			FplID:    wm.ID,
			FullName: fmt.Sprintf("%s %s", wm.FirstName, wm.LastName),
		}

		results <- am
	}
}
