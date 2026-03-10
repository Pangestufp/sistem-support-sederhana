package helper

import "time"

func FormatTimeRFC3339(t time.Time) string {
	if t.IsZero() {
		return "-"
	}
	return t.Format(time.RFC3339)
}
