package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var c = &http.Client{Timeout: time.Duration(5) * time.Second}

func Get(url string) {

	resp, err := c.Get(url)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(ioutil.Discard, resp.Body)
	fmt.Printf("status_code: %d\n", resp.StatusCode)
}
