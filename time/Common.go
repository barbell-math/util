package time

import "time"

// Gets the number of days between two dates, including negative days if before
// is not before before. ;)
func DaysBetween(before time.Time, after time.Time) int {
	if after.After(before) {
		return int(after.Sub(before).Hours() / 24)
	} else {
		return -int(before.Sub(after).Hours() / 24)
	}
}
