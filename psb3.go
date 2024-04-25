package main

import "time"

const (
	HHMMSS12h string = "03:04:05PM"
	HHMMSS24h string = "15:04:05"
)

func ConvertTo24HourFormat(s string) (string, error) {
	t, err := time.Parse(HHMMSS12h, s)
	if err != nil {
		return "", err
	}

	return t.Format(HHMMSS24h), nil
}
