package wrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Wrapper is a helper interface around FPL API
type Wrapper interface {
	GetManager(id int) (*Manager, error)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}
