package handlers

import "net/http"

// AvailableUrl checking url that is correct or not
func AvailableUrl(url string) bool {
	_, err := http.Get(url)
	if err != nil {
		return false
	}
	return true
}
