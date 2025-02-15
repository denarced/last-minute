// Package lastmin.
package lastmin

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var datePattern = regexp.MustCompile(`(\d\d\d\d).?(\d\d).?(\d\d).?(\d\d).?(\d\d).?(\d\d)`)

type DatedLine struct {
	Date time.Time
	Line string
}

func ParseDate(line string) (time.Time, error) {
	matches := datePattern.FindStringSubmatch(line)
	if matches == nil {
		return time.Time{}, fmt.Errorf("nothing found")
	}
	d := []int{}
	for i := 1; i < 7; i++ {
		value, err := strconv.Atoi(matches[i])
		if err != nil {
			return time.Time{}, err
		}
		d = append(d, value)
	}
	return time.Date(d[0], time.Month(d[1]), d[2], d[3], d[4], d[5], 0, time.Local), nil
}

func FilterLines(lines []DatedLine, earliest, now time.Time, seconds int) (filtered []DatedLine) {
	limit := now.Add(time.Duration(-seconds) * time.Second)
	for _, each := range lines {
		if each.Date.Before(limit) || each.Date.Before(earliest) {
			continue
		}
		filtered = append(filtered, each)
	}
	return
}
