package dockerhub

import (
	"encoding/json"
	"net/http"
	"time"
)

// DockerSearchImage is a struct to represent the data of interest for a docker image
type DockerSearchImage struct {
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Publisher struct {
		Name string `json:"name"`
	} `json:"publisher"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"short_description"`
	PullCount   string    `json:"pull_count"`
	Type        string    `json:"filter_type"`
}

// SearchAPIResponse is a struct to represent the top level response data from a Docker Hub search API call
type SearchAPIResponse struct {
	Count    int                 `json:"count"`
	Page     int                 `json:"page"`
	PageSize int                 `json:"page_size"`
	Images   []DockerSearchImage `json:"summaries"`
}

// Search is the API that performs a search against Docker Hub and returns the result
func Search(search string, page uint32, limit uint32) (SearchAPIResponse, error) {
	var results SearchAPIResponse

	url := SearchURL(search, page, limit)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return results, err
	}

	// The Search-Version: v3 header is needed to fetch the desired JSON response from Docker Hub API
	req.Header.Set("Search-Version", "v3")
	req.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return results, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&results)
	if err != nil {
		return results, err
	}

	return results, nil
}
