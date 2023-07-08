package service

type Create interface {
	CreateNomadConfiguration(name, path, from string) (bool, error)
}

type Service struct {
	Create Create
}

func NewService() *Service {
	return &Service{
		Create: NewCreateService(),
	}
}
