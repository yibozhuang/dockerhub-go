package dockerhub

import "fmt"

// SearchURL formulates the Docker Hub url based on input parameters
func SearchURL(search string, page uint32, limit uint32) string {
	base := "https://hub.docker.com/api/content/v1/products/search"
	return fmt.Sprintf("%s?type=image&q=%s&page=%d&page_size=%d", base, search, page, limit)
}

// InfoURL formulates the Docker Hub image info url based on namespace and image name
func InfoURL(namespace string, name string) string {
	base := "https://hub.docker.com/v2/repositories"
	return fmt.Sprintf("%s/%s/%s", base, namespace, name)
}

// TagsURL formulates the Docker Hub image tags url based on namespace and image name
func TagsURL(namespace string, name string, page uint32, limit uint32) string {
	base := "https://hub.docker.com/v2/repositories"
	return fmt.Sprintf("%s/%s/%s/tags/?page_size=%d&page=%d", base, namespace, name, limit, page)
}
