package dockerhub

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// DockerTagDetail is a struct to represent the details of a particular tag
type DockerTagDetail struct {
	Architecture string `json:"architecture"`
	Hash         string `json:"digest"`
	Os           string `json:"os"`
	Size         int    `json:"size"`
}

// DockerTag is a struct to describe the high level info of a docker tag
type DockerTag struct {
	Name        string            `json:"name"`
	LastUpdated time.Time         `json:"last_updated"`
	Details     []DockerTagDetail `json:"images"`
}

// TagAPIResponse is a struct that stores the tags from the API request made to Docker Hub
type TagAPIResponse struct {
	Count int         `json:"count"`
	Tags  []DockerTag `json:"results"`
}

// Tags is the API that queries the tags for a particular image from Docker Hub
func Tags(image string, page uint32, limit uint32) (TagAPIResponse, error) {
	var results TagAPIResponse

	namespace := "library"
	name := image

	splitted := strings.Split(image, "/")
	if len(splitted) > 1 {
		namespace = splitted[0]
		name = splitted[1]
	}

	url := TagsURL(namespace, name, page, limit)

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
