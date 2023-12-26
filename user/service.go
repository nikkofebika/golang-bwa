package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	GetAll() ([]User, error)
	RegisterUser(input RegisterUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}

	// var user User = User{
	// 	Name:     input.Name,
	// 	Email:    input.Email,
	// 	Password: string(password),
	// 	Role:     "user",
	// }

	var user User = User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Password = string(password)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
