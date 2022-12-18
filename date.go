package date

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrDateInvalid = errors.New("date invalid")

	// Custom date formats
	// Default is the first
	dateFormats = []string{
		// YYYY->MM->DD
		"2006-01-02",
		"2006/01/02",
		"2006.01.02",
		"2006 01 02",
		"20060102",
		// DD->MM->YYYY
		"02-01-2006",
		"02/01/2006",
		"02.01.2006",
		"02 01 2006",
		"02012006",
	}
	defaultDateFormat = dateFormats[0]

	dateToday     = "today"
	dateYesterday = "yesterday"

	dateTimeFormats = []string{
		// YYYY->MM->DD hh->mm->ss
		"2006-01-02 15:04:05",
	}
	defaultDateTimeFormat = dateTimeFormats[0]
)

// Convert YYYY/MM/DD and more to RFC3339
func ToRFC3339(date string, location *time.Location) (string, error) {
	// UTC should be default
	if location == nil {
		location = time.UTC
	}

	if strings.EqualFold(date, dateToday) {
		return time.Now().In(location).Format(time.RFC3339), nil
	}

	if strings.EqualFold(date, dateYesterday) {
		return time.Now().Add(-time.Hour * 24).In(location).Format(time.RFC3339), nil
	}

	// Try to parse date from custom formats
	var t time.Time
	for _, dateFormat := range dateFormats {
		var err error
		t, err = time.ParseInLocation(dateFormat, date, location)
		if err != nil {
			continue
		}

		break
	}

	if t.IsZero() {
		return "", ErrDateInvalid
	}

	return t.Format(time.RFC3339), nil
}

// Convert RFC3339 to default custom date format
func FromRFC3339(rfc3339 string, location *time.Location) (string, error) {
	t, err := time.ParseInLocation(time.RFC3339, rfc3339, location)
	if err != nil {
		return "", fmt.Errorf("failed to parse time: %w", err)
	}

	return t.Format(defaultDateFormat), nil
}

// Whatever date format is, always return default date format
func ToDefaultDate(date string, location *time.Location) (string, error) {
	dateInRFC3339, err := ToRFC3339(date, location)
	if err != nil {
		return "", err
	}

	return FromRFC3339(dateInRFC3339, location)
}

// Return all support date formats
// For showing only
func SupportDateFormats() string {
	result := append([]string{}, dateFormats...)
	result = append(result, dateToday, dateYesterday)

	return strings.Join(result, ", ")
}

// Format time to default date format
func FormatDateByDefault(t time.Time, location *time.Location) string {
	return t.In(location).Format(defaultDateFormat)
}

// Format time to default date time format
func FormatDateTimeByDefault(t time.Time, location *time.Location) string {
	return t.In(location).Format(defaultDateTimeFormat)
}
