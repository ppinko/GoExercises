package main

import (
	"fmt"
	"time"
)

// https://exercism.org/tracks/go/exercises/booking-up-for-beauty

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	timeObj := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", timeObj.Weekday(), timeObj.Month(), timeObj.Day(), timeObj.Year(), timeObj.Hour(), timeObj.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	now := time.Now()
	date := time.Date(now.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
	return date
}

func main() {
	fmt.Println()

	fmt.Println(Schedule("7/25/2019 13:45:00")) // 2019-07-25 13:45:00 +0000 UTC

	fmt.Println(HasPassed("July 25, 2019 13:45:00")) // Output: true

	fmt.Println(IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")) // Output: true

	fmt.Println(Description("7/25/2019 13:45:00")) // Output: "You have an appointment on Thursday, July 25, 2019, at 13:45."
	fmt.Println(Description("6/6/2005 10:30:00"))  // Output: "You have an appointment on Monday, June 6, 2005, at 10:30."

	fmt.Println(AnniversaryDate()) // Output: 2020-09-15 00:00:00 +0000 UTC
}
