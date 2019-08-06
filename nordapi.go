// Package nordapi provides functions for accessing the NordAPI public API
// it is thread safe.
package nordapi

import (
	"encoding/json"
	"net/http"
	"time"
)

// client allows the http client to persist between api calls as per net/http's guidelines
var client = &http.Client{
	Timeout: time.Second * 10,
}

func getAndUnmarshall(url string, obj interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(resp.Body)
	return dec.Decode(obj)
}
