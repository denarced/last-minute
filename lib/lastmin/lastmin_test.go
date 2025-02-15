package lastmin

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseDate(t *testing.T) {
	run := func(name, line string, expected time.Time, expectedErrorMessage string) {
		t.Run(name, func(t *testing.T) {
			req := require.New(t)
			actual, err := ParseDate(line)
			if expectedErrorMessage == "" {
				req.Equal(expected, actual)
			} else {
				req.Errorf(err, expectedErrorMessage)
			}
		})
	}

	run("Full ISO8601", "My kind of 2024-01-01T00:00:00Z :)", createDate(2024, 1, 1), "")
	run(
		"No separators",
		"19991231235959",
		time.Date(1999, time.December, 31, 23, 59, 59, 0, time.Local),
		"",
	)
	run("Empty", "", time.Time{}, "nothing found")
}

func createDate(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func TestFilterLines(t *testing.T) {
	now := createDate(1600, 6, 6)
	minutes := 3
	// First "|": Limit defined by "minutes".
	// Second "|": limit defined by "now" in the real command, when it was started.
	// Third "|": now, as in when this specific filtering is done.
	// "_": timeline without anything (i.e. empty).
	// "-": position of the timestamp on the timeline.
	// "+": combined "|" and "-".
	lines := []DatedLine{
		{Line: "_-||_|_", Date: now.Add(time.Duration(-minutes-1) * time.Minute)},
		{Line: "__+|_|_", Date: now.Add(time.Duration(-minutes) * time.Minute)},
		{Line: "__|+_|_", Date: now.Add(time.Duration(-minutes+1) * time.Minute)},
		{Line: "__||-|_", Date: now.Add(time.Duration(-1) * time.Minute)},
		{Line: "__||_+_", Date: now},
		{Line: "__||_|-", Date: now.Add(time.Duration(1) * time.Minute)},
	}
	initialNow := now.Add(time.Duration(-minutes+1) * time.Minute)
	filtered := FilterLines(lines, initialNow, now, minutes*60)
	require.Equal(t, lines[2:], filtered)
}
