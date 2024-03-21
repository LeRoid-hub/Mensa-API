package fetch

import (
	"net/http"
	"net/url"
)

func Fetch(path string) (*http.Response, error) {
	baseurl := "https://www.imensa.de/"
	queryurl := baseurl + "/" + path
	u, err := url.ParseRequestURI(queryurl)
	if err != nil {
		return nil, err
	}
	return http.Get(u.String())
}
