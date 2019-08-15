// Package nordapi provides functions for accessing the NordAPI public API
// it is thread safe.
package nordapi

import (
	"encoding/json"
	"fmt"
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
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("getting \"%v\"; %v", url, resp.Status)
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(obj)
	if err != nil {
		return fmt.Errorf("decoding \"%v\"; %v", url, err)
	}
	return nil
}
