package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Gender = input.Gender

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	CheckEmail, _ := s.repository.FindByEmail(user.Email)
	if CheckEmail.ID > 0 {
		return CheckEmail, errors.New("The email you are using is already registered")
	}


	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	userData, err := s.repository.FindByEmail(newUser.Email)
	if err != nil {
		return user, err
	}

	return userData, nil

}


func (s *service) Login(input LoginInput) (User, error){
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, errors.New("Password does not match")
	}

	return user, nil


}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error){
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil

}

	func (s *service) SaveAvatar(ID int, fileLocation string) (User, error){
		// find user by id
		user, err := s.repository.FindByID(ID)
		if err != nil {
			return user, err
		}

		// update atribute file_avatar 
		user.FileAvatar = fileLocation

		// save field file_avatar
		updateUser, err := s.repository.Update(user)
		if err != nil {
			return updateUser, err
		}

		return updateUser, nil
	}


// mapping struct input ke strcut user
// simpan struct User melalui repository
