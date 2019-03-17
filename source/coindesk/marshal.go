package coindesk

import "time"

// https://www.coindesk.com/api
type response struct {
	Time struct {
		UpdatedISO time.Time `json:"updatedISO"`
	} `json:"time"`
	Bpi struct {
		USD struct {
			Ratefloat float64 `json:"rate_float"`
		} `json:"USD"`
	} `json:"bpi"`
}
