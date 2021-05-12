package adding

type service struct {
	r Repository
}

type FieldError struct {
	Field string
}

func (s *service) Add(name string) (*Project, error) {
	return s.r.Add(name)
}

type Repository interface {
	Add(name string) (*Project, error)
}

func NewService(r Repository) *service {
	return &service{r}
}
