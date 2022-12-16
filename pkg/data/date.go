package data

import (
	"time"
)

var dateFormats []string = []string{
	time.Layout,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
}

type Date struct {
	ts time.Time
}

func (x *Date) String() string {
	return x.ts.Format("2006-01-02T15:04:05")
}

func (x *Date) Before(d time.Time) bool {
	return x.ts.Before(d)
}

func ParseDate(raw string) Date {
	return Date{ts: parseDate(raw)}
}

func parseDate(raw string) time.Time {
	for _, f := range dateFormats {
		if t, err := time.Parse(f, raw); err == nil {
			return t
		}
	}
	return time.Now()
}
