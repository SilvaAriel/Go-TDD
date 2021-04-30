package listing

type Repository interface {
	ListAllProjects() []Project
}

type Service struct {
	r Repository
}

func (s *Service) ListAllProjects() []Project {
	return s.r.ListAllProjects()
}

func NewService(r Repository) *Service {
	return &Service{r}
}
