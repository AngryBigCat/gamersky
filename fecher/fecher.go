package fecher

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

var Total int

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

func GetTotal() {
	Total = 123;
}