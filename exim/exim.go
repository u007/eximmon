package exim

import "time"

// ParseDate get local time of a date
func ParseDate(thedate string) (time.Time, error) {
	local := time.Now()

	if len(thedate) > 10 {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", thedate, local.Location())
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	} else {
		t, err := time.ParseInLocation("2006-01-02", thedate, local.Location())
		if err != nil {
			return time.Time{}, err
		}
		return t, nil
	}
}
