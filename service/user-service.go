package service

import (
	"lab-go-gin-jwt/dto"
	"lab-go-gin-jwt/entity"
	"lab-go-gin-jwt/repository"
	"log"

	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

// Update implements UserService
func (s *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Panicf("Failed map %v:", err)
	}
	updatedUser := s.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

// Profile implements UserService
func (s *userService) Profile(userID string) entity.User {
	return s.userRepository.ProfileUser(userID)
}
