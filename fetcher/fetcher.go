package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var current = 1

func Get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return ioutil.ReadAll(res.Body)
}
