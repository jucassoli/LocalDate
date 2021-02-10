package localdate

import "time"

func Format(gientime time.Time) string {
	return gientime.Format("dd-MM-yyyy")
}
