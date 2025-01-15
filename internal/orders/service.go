package orders

type Service interface {
	GetAll(requestFields []string) (*[]Orders, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) GetAll(requestFields []string) (*[]Orders, error) {
	data, err := s.repo.GetAll(requestFields)
	if err != nil {
		return nil, err
	}
	return data, nil
}
