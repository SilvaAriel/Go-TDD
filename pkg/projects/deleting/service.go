package deleting

type Repository interface {
	Delete(id int) (*Project, error)
}

type service struct {
	r Repository
}

func (s *service) Delete(id int) (*Project, error) {
	return s.r.Delete(id)
}

func NewService(r Repository) *service {
	return &service{r}
}
