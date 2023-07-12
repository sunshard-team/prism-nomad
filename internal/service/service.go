package service

type Output interface {
	CreateNomadConfiguration(name, path, from string) (bool, error)
}

type Service struct {
	Output Output
}

func NewService() *Service {
	return &Service{
		Output: NewOutputService(),
	}
}
