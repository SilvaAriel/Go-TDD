package deleting

import "time"

type Report struct {
	ID        int
	name      string
	CreatedAt time.Time
	Success   bool
}
