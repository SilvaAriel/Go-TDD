package updating

type service struct {
	r Repository
}

func (s *service) Update(id int, name string) (*Project, error) {
	updatedValue, err := s.r.Update(id, name)
	return updatedValue, err
}

func NewService(r Repository) *service {
	return &service{r}
}

type Repository interface {
	Update(id int, name string) (*Project, error)
}
