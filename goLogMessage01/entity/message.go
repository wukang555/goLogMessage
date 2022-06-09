package entity

import "time"

type Message struct {
	FileName string
	Body     string
	GetTime  time.Time
}
