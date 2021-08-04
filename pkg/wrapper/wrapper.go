package wrapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var errHTTPStatusNotOK error = errors.New("Returned http status is not 200")
var errReadFailure error = errors.New("Failed to read the response")
var errUnmarshalFailure error = errors.New("Failed to unmarshal data")

// Wrapper is a helper interface around FPL API
type Wrapper interface {
	GetManager(id int) (*Manager, error)
	GetManagersCount() (int, error)
}

type wrapper struct {
	client  *http.Client
	baseURL string
}

// NewWrapper returns instance of an FPL API wrapper
func NewWrapper() Wrapper {
	return &wrapper{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		baseURL: "https://fantasy.premierleague.com/api",
	}
}

// GetManager returns data from FPL API "/api/entry/{managerID}/" endpoint
func (w *wrapper) GetManager(id int) (*Manager, error) {
	url := fmt.Sprintf(w.baseURL+"/entry/%d/", id)
	var manager Manager

	err := w.fetchData(url, &manager)
	if err != nil {
		return nil, err
	}

	return &manager, nil
}

// GetManagersCount returns number of total managers, from FPL API "/api/bootstrap-static/" endpoint
func (w *wrapper) GetManagersCount() (int, error) {
	url := fmt.Sprintf(w.baseURL + "/bootstrap-static/")
	var tp totalPlayers

	err := w.fetchData(url, &tp)
	if err != nil {
		return 0, err
	}

	return tp.Count, nil
}

func (w *wrapper) fetchData(url string, data interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "app")

	resp, err := w.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// TODO check rate limiting error
	if resp.StatusCode != http.StatusOK {
		return errHTTPStatusNotOK
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errReadFailure
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return errUnmarshalFailure
	}

	return nil
}
