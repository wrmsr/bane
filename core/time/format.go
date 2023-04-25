package time

import (
	"errors"
	"time"
)

//

const (
	RFC3339NZ = "2006-01-02T15:04:05"
)

func ParseAny(layouts []string, s string) (time.Time, error) {
	for _, l := range layouts {
		if t, err := time.Parse(l, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, errors.New("failed to parse time")
}

//

const (
	ANSICDate    = "Mon Jan _2"
	RubyDateDate = "Mon Jan 02"
	RFC822Date   = "02 Jan 06"
	RFC850Date   = "Monday, 02-Jan"
	RFC1123Date  = "Mon, 02 Jan 2006"
	RFC3339Date  = "2006-01-02"

	ANSICTime       = "15:04:05 2006"
	UnixDateTime    = "15:04:05 MST 2006"
	RubyDateTime    = "15:04:05 -0700 2006"
	RFC822Time      = "15:04 MST"
	RFC822ZTime     = "15:04 -0700" // RFC822 with numeric zone
	RFC850Time      = "15:04:05 MST"
	RFC1123ZTime    = "15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339Time     = "15:04:05Z07:00"
	RFC3339NanoTime = "15:04:05.999999999Z07:00"
)

var dateLayouts = map[string]string{
	time.ANSIC:       ANSICDate,
	time.UnixDate:    ANSICDate,
	time.RubyDate:    RubyDateDate,
	time.RFC822:      RFC822Date,
	time.RFC822Z:     RFC822Date,
	time.RFC850:      RFC850Date,
	time.RFC1123:     RFC1123Date,
	time.RFC1123Z:    RFC1123Date,
	time.RFC3339:     RFC3339Date,
	time.RFC3339Nano: RFC3339Date,
}

var timeLayouts = map[string]string{
	time.ANSIC:       ANSICTime,
	time.UnixDate:    UnixDateTime,
	time.RubyDate:    RubyDateTime,
	time.RFC822:      RFC822Time,
	time.RFC822Z:     RFC822ZTime,
	time.RFC850:      RFC850Time,
	time.RFC1123:     RFC850Time,
	time.RFC1123Z:    RFC1123ZTime,
	time.RFC3339:     RFC3339Time,
	time.RFC3339Nano: RFC3339NanoTime,
}
