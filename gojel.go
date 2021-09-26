package gojel

import (
	"time"
)

func SetTimeout(callback func(), millis int64) *time.Timer {
	timer := time.NewTimer(time.Duration(millis * int64(time.Millisecond)))
	go func() {
		<-timer.C
		callback()
	}()
	return timer
}

func SetInterval(callback func(), millis int64) *time.Ticker {
	ticker := time.NewTicker(time.Duration(millis * int64(time.Millisecond)))
	go func() {
		<-ticker.C
		callback()
	}()
	return ticker
}
