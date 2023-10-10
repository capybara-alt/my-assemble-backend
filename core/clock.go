package core

import "time"

var _time time.Time

func SetFakeTime(t time.Time) {
	_time = t
}

func ResetFakeTime() {
	_time = time.Time{}
}

func Now() time.Time {
	if !_time.IsZero() {
		return _time
	}

	return time.Now()
}
