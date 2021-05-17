package listing

type Repository interface {
	GetById(projectId int, reportId int) ([]Report, error)
	GetAll(projectId int) ([]Report, error)
}

type service struct {
	r Repository
}

func (s *service) GetById(projectId int, reportId int) ([]Report, error) {
	return s.r.GetById(projectId, reportId)
}

func (s *service) GetAll(projectId int) ([]Report, error) {
	return s.r.GetAll(projectId)
}

func NewService(r Repository) *service {
	return &service{r}
}
