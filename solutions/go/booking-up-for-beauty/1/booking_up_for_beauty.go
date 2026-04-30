package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
    t, _ := time.Parse(layout,date)
    return t
	panic("Please implement the Schedule function")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"

	t, _ := time.Parse(layout, date)
	return time.Now().After(t)
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
    if err != nil {
        return false
    }
    hour := t.Hour()
    return hour >= 12 && hour < 18
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    t, _ := time.Parse("1/2/2006 15:04:05", date)

    return t.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().Year()
    return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	panic("Please implement the AnniversaryDate function")
}
