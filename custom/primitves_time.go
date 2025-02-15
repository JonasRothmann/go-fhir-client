package custom

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	dateFormat = "2006-01-02" // layout for date-only
	timeFormat = "15:04:05"   // layout for time-only
)

// Date wraps time.Time but is used for date-only values.
type Date struct {
	time.Time
}

// UnmarshalJSON parses a JSON string in "YYYY-MM-DD" format.
func (d *Date) UnmarshalJSON(b []byte) error {
	// Remove surrounding quotes.
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		d.Time = time.Time{}
		return nil
	}
	parsed, err := time.Parse(dateFormat, s)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

// MarshalJSON outputs the date in "YYYY-MM-DD" format.
func (d Date) MarshalJSON() ([]byte, error) {
	formatted := d.Format(dateFormat)
	return []byte(fmt.Sprintf("\"%s\"", formatted)), nil
}

// TimeOfDay wraps time.Time but is used for time-only values.
type Time struct {
	time.Time
}

// UnmarshalJSON parses a JSON string in "HH:MM:SS" format.
func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		t.Time = time.Time{}
		return nil
	}
	parsed, err := time.Parse(timeFormat, s)
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}

// MarshalJSON outputs the time in "HH:MM:SS" format.
func (t Time) MarshalJSON() ([]byte, error) {
	formatted := t.Format(timeFormat)
	return []byte(fmt.Sprintf("\"%s\"", formatted)), nil
}

type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	// Try to parse using various layouts:
	layouts := []string{
		"2006-01-02T15:04:05Z07:00", // full RFC3339
		"2006-01-02",                // date only
		"2006-01",                   // year-month
		"2006",                      // year only
	}
	var err error
	for _, layout := range layouts {
		var t time.Time
		if t, err = time.Parse(layout, s); err == nil {
			dt.Time = t
			return nil
		}
	}
	return fmt.Errorf("unable to parse FHIRTime: %s", s)
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	// Always output full RFC3339 for consistency.
	return json.Marshal(dt.Time.Format("2006-01-02T15:04:05Z07:00"))
}
