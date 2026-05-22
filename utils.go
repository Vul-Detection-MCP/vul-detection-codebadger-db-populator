package main

import "time"

const MainTimeFormat = time.RFC3339Nano

func GetStableTime() time.Time {
	return time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
}

func ToTimeFormat(given time.Time) string {
	return given.Format(time.RFC3339Nano)
}
