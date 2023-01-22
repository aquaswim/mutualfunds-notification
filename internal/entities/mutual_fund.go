package entities

import "time"

type MutualFund struct {
	Name       string
	LatestDate string
	NavHistory []NavLog
}

type NavLog struct {
	Date  time.Time
	Value float32
}
