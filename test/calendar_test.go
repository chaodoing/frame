package test

import (
	`testing`
	`time`
	
	`github.com/chaodoing/frame/calendar`
)

func TestCalendars(t *testing.T) {
	// start := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
	// finish := time.Date(time.Now().Year(), time.Now().Month()+1, 0, 0, 0, 0, 0, time.Local)
	// t.Log(start, finish)
	
	dates := calendar.Calendar(time.Now().Year(), time.Now().Month())
	t.Log(dates)
}
