// Package nordapi provides functions for accessing the NordAPI public API
// it is thread safe.
package nordapi

import (
	"net/http"
	"time"
)

// client allows the http client to persist between api calls as per net/http's guidelines
var client = &http.Client{
	Timeout: time.Second * 10,
}
