package fetch

import (
	"net/http"
)

func Fetch(url string) (*http.Response, error) {
	u, err := url.ParseRequestURI(url)
	if err != nil {
		return nil, err
	}

	return http.Get(url)
}
