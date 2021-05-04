package adding

import (
	"fmt"
)

type service struct {
	r Repository
}

type FieldError struct {
	Field string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%q is required", e.Field)
}

func (e *FieldError) Is(target error) bool {
	t, ok := target.(*FieldError)
	if !ok {
		return false
	}
	return e.Field == t.Field
}

func (s *service) Add(name string) (*Project, error) {
	if name != "" {
		return s.r.Add(name)
	}
	return nil, &FieldError{Field: "name"}
}

type Repository interface {
	Add(name string) (*Project, error)
}

func NewService(r Repository) *service {
	return &service{r}
}
