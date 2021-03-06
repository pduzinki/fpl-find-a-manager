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
	15 * time.Second,
	20 * time.Second,
	15 * time.Minute,
	20 * time.Minute,
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

	for {
		var totalManagers int
		var addedManagers int
		var err error

		for {
			totalManagers, err = mc.w.GetManagersCount()
			if err != nil {
				retryAfter := 5 * time.Minute
				log.Println("Failed to retrieve number of FPL managers, retry in", retryAfter)
				time.Sleep(retryAfter)
				continue
			}
			log.Println("FPL managers in the game:", totalManagers)
			break
		}

		addedManagers, err = mc.ms.ManagersCount()
		if err != nil {
			log.Println("Failed to retrieve number of FPL managers in the database!")
			return
		}
		log.Println("FPL managers in the database:", addedManagers)

		var goroutinesCount = 32
		var numJobs = 1000
		jobs := make(chan int, numJobs)
		results := make(chan models.Manager, numJobs)

		for w := 1; w <= goroutinesCount; w++ {
			go mc.getManagersFromFPL(w, jobs, results)
		}

		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		tickerDone := make(chan bool)
		go func() {
			for {
				select {
				case <-tickerDone:
					return
				case <-ticker.C:
					addedManagers, err = mc.ms.ManagersCount()
					if err != nil {
						log.Println("Failed to retrieve number of FPL managers in the database!")
						continue
					}
					log.Printf("Managers in the database: %v. Coverage: %.2f%%\n",
						addedManagers, 100*float64(addedManagers)/float64(totalManagers))
				}
			}
		}()

		for totalManagers > addedManagers {
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

			sort.Sort(models.Managers(managers)) // so ID == fplID
			mc.ms.AddManagers(managers)

			addedManagers += numJobs
		}

		tickerDone <- true
		close(jobs)

		// all managers added, sleep and then add newcomers
		log.Println("Current FPL managers added, going to sleep for an hour now")
		time.Sleep(1 * time.Hour)
	}
}

func (mc *ManagerController) getManagersFromFPL(id int, jobs chan int, results chan<- models.Manager) {
	w := wrapper.NewWrapper()

	for j := range jobs {
		wm, err := w.GetManager(j)
		if err == wrapper.ErrHTTPStatusNotFound {
			// manager not found, let's add it as it is to the db, and keep going
			log.Printf("FPL API call returned http 404, worker %v. Job: %v", id, j)

			am := models.Manager{
				FplID:    j,
				FullName: "Manager Not Found",
			}
			results <- am
			continue
		} else if err == wrapper.ErrHTTPTooManyRequests {
			// hit rate limit, let's sleep here for a bit, return job to pool
			jobs <- j
			sleepDuration := sleeps[rand.Intn(len(sleeps))]
			log.Printf("FPL API call returned http 429, worker %v going to sleep for %v. Job: %v", id, sleepDuration, j)
			time.Sleep(sleepDuration)
			continue
		} else if err != nil {
			jobs <- j
			log.Printf("FPL API call returned '%v', worker %v going to sleep. Job: %v", err, id, j)
			time.Sleep(sleeps[len(sleeps)-1])
			continue
		}

		am := models.Manager{
			FplID:    wm.ID,
			FullName: fmt.Sprintf("%s %s", wm.FirstName, wm.LastName),
		}

		results <- am
	}
}
