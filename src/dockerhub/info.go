package dockerhub

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// DockerInfoImage is a struct to represent the data of interest from the docker image
type DockerInfoImage struct {
	Name            string    `json:"name"`
	NameSpace       string    `json:"namespace"`
	Type            string    `json:"repository_type"`
	Description     string    `json:"description"`
	PullCount       int       `json:"pull_count"`
	FullDescription string    `json:"full_description"`
	LastUpdated     time.Time `json:"last_updated"`
}

// Info is the API that fetches the detailed info of a given image from Docker Hub
func Info(image string) (DockerInfoImage, error) {
	var results DockerInfoImage

	namespace := "library"
	name := image

	splitted := strings.Split(image, "/")
	if len(splitted) > 1 {
		namespace = splitted[0]
		name = splitted[1]
	}

	url := InfoURL(namespace, name)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return results, err
	}

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
