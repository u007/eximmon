package exim

import "time"

func ParseDate(thedate string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", thedate)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
