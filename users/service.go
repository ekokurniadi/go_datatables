package users

type Service interface {
	GetAllUser(input DTJson) ([]User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUser(input DTJson) ([]User, error) {
	inputUser := DTJson{}
	inputUser.Draw = input.Draw
	inputUser.Columms = input.Columms
	inputUser.Length = input.Length
	inputUser.Orders = input.Orders
	inputUser.Start = input.Start
	inputUser.Search.Value = input.Search.Value
	users, err := s.repository.GetUsers(inputUser)
	if err != nil {
		return users, err
	}
	return users, nil
}
