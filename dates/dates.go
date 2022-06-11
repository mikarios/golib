package dates

import (
	"time"
)

func MatchesDay(now, check time.Time) bool {
	ny, nm, nd := now.Date()
	cy, cm, cd := check.Date()

	return ny == cy && nm == cm && nd == cd
}
