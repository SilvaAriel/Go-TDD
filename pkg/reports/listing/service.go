package listing

type Repository interface {
	GetById(projectId int, reportId int) (*Report, error)
}

type service struct {
	r Repository
}

func (s *service) GetById(projectId int, reportId int) (*Report, error) {
	return s.r.GetById(projectId, reportId)
}

func NewService(r Repository) *service {
	return &service{r}
}
