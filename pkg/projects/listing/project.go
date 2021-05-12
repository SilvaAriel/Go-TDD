package listing

import "time"

type Project struct {
	ID        int
	Name      string
	Status    Status
	CreatedAt time.Time
	Reports   []Report
}
