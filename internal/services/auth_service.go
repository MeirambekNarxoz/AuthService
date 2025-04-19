package service

import (
	"authService/internal/models"
	"authService/internal/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	jwtKey   []byte
}

func NewAuthService(userRepo *repository.UserRepository, jwtKey string) *AuthService {
	return &AuthService{userRepo: userRepo, jwtKey: []byte(jwtKey)}
}

// Register registers a new user with default RoleID = 1
func (uc *AuthService) Register(username, password string) (string, error) {
	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
		RoleID:   1,
	}
	err = uc.userRepo.CreateUser(user)
	if err != nil {
		return "", errors.New("failed to create user")
	}

	// Генерация JWT-токена
	token, err := generateJwtToken(user, uc.jwtKey)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

// Login authenticates a user and returns a JWT token
func (uc *AuthService) Login(username, password string) (string, error) {
	// Поиск пользователя по имени
	user, err := uc.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Генерация JWT-токена
	token, err := generateJwtToken(user, uc.jwtKey)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

// generateJwtToken generates a JWT token for the given user
func generateJwtToken(user *models.User, jwtKey []byte) (string, error) {
	// Создание токена с данными пользователя
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"username":  user.Username,
		"user_role": user.Role.Name,                        // Статическое значение роли
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // Токен действителен 24 часа
	})

	// Подписание токена
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return signedToken, nil
}
