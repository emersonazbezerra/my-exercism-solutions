package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	afterNoonStart, _ := time.Parse("15:04", "12:00")
	afterNoonEnd, _ := time.Parse("15:04", "18:00")

	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

	return t.Hour() >= afterNoonStart.Hour() && t.Hour() < afterNoonEnd.Hour()
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return fmt.Sprintf(
		"You have an appointment on %s, %s %d, %d, at %s.",
		t.Weekday(), t.Month(), t.Day(), t.Year(), t.Format("15:04"),
	)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
