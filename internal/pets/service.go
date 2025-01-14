package pets

type Service interface {
	GetAll() (*[]Pets, error)
	GetOne(id int32) (*Pets, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s service) GetAll() (*[]Pets, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s service) GetOne(id int32) (*Pets, error) {
	pet, err := s.repo.GetOne(id)
	if err != nil {
		return nil, err
	}

	return pet, nil
}
