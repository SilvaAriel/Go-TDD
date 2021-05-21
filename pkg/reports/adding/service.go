package adding

type Repository interface {
	Add(projectId int, name string, success bool) (*Report, error)
}

type service struct {
	r Repository
}

func (s *service) Add(projectId int, name string, success bool) (*Report, error) {
	return s.r.Add(projectId, name, success)
}

func NewService(r Repository) *service {
	return &service{r}
}
