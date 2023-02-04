package routes

import "time"

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL            string `json:"url"`
	CustomShort    string `json:"short"`
	Expiry         string `json:"expiry"`
	XRateRemaining int    `json:"rate_limit"`
	XRateLimitRest int    `json:"rate_limit_remaining"`
}
