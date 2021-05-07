package listing

type Repository interface {
	ListAll() []Project
	GetById(id int) ([]Project, error)
}

type service struct {
	r Repository
}

func (s *service) ListAll() []Project {
	return s.r.ListAll()
}

func (s *service) GetById(id int) ([]Project, error) {
	return s.r.GetById(id)
}

func NewService(r Repository) *service {
	return &service{r}
}
