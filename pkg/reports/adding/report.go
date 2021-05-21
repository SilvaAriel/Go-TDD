package adding

import "time"

type Report struct {
	ID        int
	ProjectID int
	name      string
	CreatedAt time.Time
	Success   bool
}
