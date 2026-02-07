package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/repository"
	"github.com/Albaihaqi354/FinalPhase3.git/pkg"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	authRepository *repository.AuthRepository
	redis          *redis.Client
}

func NewAuthService(authRepository *repository.AuthRepository, rdb *redis.Client) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		redis:          rdb,
	}
}

func (a AuthService) Register(ctx context.Context, newUser dto.RegisterRequest) (dto.RegisterResponse, error) {
	hc := pkg.HashConfig{}
	hc.UseRecomended()

	hp, err := hc.GenHash(newUser.Password)
	if err != nil {
		log.Println(err.Error())
		return dto.RegisterResponse{}, err
	}

	data, err := a.authRepository.CreateNewUser(ctx, newUser, hp)
	if err != nil {
		log.Println(err.Error())
		return dto.RegisterResponse{}, err
	}

	response := dto.RegisterResponse{
		Id:    data.Id,
		Email: data.Email,
	}
	return response, nil
}

func (a AuthService) Login(ctx context.Context, loginReq dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := a.authRepository.FindUserByEmail(ctx, loginReq.Email)
	if err != nil {
		log.Println(err.Error())
		return dto.LoginResponse{}, errors.New("Invalid Email Or Password")
	}

	hc := pkg.HashConfig{}
	hc.UseRecomended()

	isValid, err := hc.ComparePwdAndHash(loginReq.Password, user.Password)
	if err != nil {
		log.Println(err.Error())
		return dto.LoginResponse{}, errors.New("Invalid Email Or Password")
	}

	if !isValid {
		return dto.LoginResponse{}, errors.New("Invalid Email Or Password")
	}

	jwtClaims := pkg.NewJWTClaims(user.Id, user.Email)
	token, err := jwtClaims.GetToken()
	if err != nil {
		log.Println(err.Error())
		return dto.LoginResponse{}, errors.New("Internal Server Error")
	}

	err = a.authRepository.SaveToken(ctx, token, time.Hour)
	if err != nil {
		log.Println("Error Save Token:", err.Error())
		return dto.LoginResponse{}, errors.New("Internal Server Error")
	}

	response := dto.LoginResponse{
		Id:    user.Id,
		Email: user.Email,
		Token: token,
	}

	return response, nil
}

func (a AuthService) Logout(ctx context.Context, token string) error {
	return a.authRepository.DeleteToken(ctx, token)
}
