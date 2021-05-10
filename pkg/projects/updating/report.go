package updating

import "time"

type Report struct {
	ID        int
	Name      string
	Success   bool
	CreatedAt time.Time
}
