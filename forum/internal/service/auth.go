package service

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/internal/storage"
	"forum/models"
	"math/rand"
	"net/smtp"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound      = errors.New("user does not exist or password incorrect")
	ErrInvalidUserName   = errors.New("invalid username")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrPasswordDontMatch = errors.New("password didn't match")
	ErrUserExist         = errors.New("user exist")
	ErrOtpCodeNotExists  = errors.New("otp code not exists")
	ErrInvlidOtpCode     = errors.New("invalid otp code")
)

type Auth interface {
	CreateUser(user models.User) error
	GenerateSessionToken(login string) (string, error)
	SignIn(login, password string) (string, error)
	ApproveOtpCode(login, code string) (string, error)
	ParseSessionToken(token string) (models.User, error)
	DeleteSessionToken(token string) error
	CheckToken(token string) error
	GetUserByToken(token string) (models.User, error)
}

type AuthService struct {
	storage storage.Auth
	cache   map[string]string
}

func newAuthService(storage storage.Auth) *AuthService {
	return &AuthService{
		storage: storage,
		cache:   make(map[string]string),
	}
}

func (s *AuthService) SignIn(login, password string) (string, error) {
	user, err := s.storage.GetUserByLogin(login)
	if err != nil {
		return "", fmt.Errorf("service: sign in: %w", err)
	}

	if err := compareHashAndPassword(user.Password, password); err != nil {
		return "", fmt.Errorf("service: sign in: %w", err)
	}

	err = s.sendOtp(user)
	if err != nil {
		return "", fmt.Errorf("service: sign in: %w", err)
	}

	return user.Email, nil
}

func (s *AuthService) sendOtp(user models.User) error {
	senderEmail := os.Getenv("OTP_SENDER_EMAIL")
	senderPassword := os.Getenv("OTP_SENDER_PASSWORD")

	recieverEmail := user.Email

	max := 10000
	min := 1000

	randonOtpCode := strconv.Itoa(rand.Intn(max-min) + min)

	subject := "OTP"
	body := randonOtpCode

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	message := []byte("To: " + recieverEmail + "\r\n" + "Subject: " + subject + "\r\n\r\n" + body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, strings.Split(recieverEmail, ","), message)
	if err != nil {
		return fmt.Errorf("service: sendOtp: %w", err)
	}

	s.cache[user.Username] = randonOtpCode

	return nil
}

func (s *AuthService) ApproveOtpCode(login, code string) (string, error) {
	user, err := s.storage.GetUserByLogin(login)
	if err != nil {
		return "", fmt.Errorf("service: approve otp code: %w", err)
	}

	originOtpCode, ok := s.cache[user.Username]
	if !ok {
		return "", fmt.Errorf("service: approve otp code: %w", ErrOtpCodeNotExists)
	}

	if originOtpCode != code {
		if err := s.sendOtp(user); err != nil {
			return "", fmt.Errorf("service: approve otp code: %w", err)
		}

		return "", fmt.Errorf("service: approve otp code: %w: sending new otp", ErrInvlidOtpCode)
	}

	token, err := s.GenerateSessionToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("service: approve otp code: %w", err)
	}

	delete(s.cache, user.Username)

	return token, nil
}

func (s *AuthService) CreateUser(user models.User) error {
	var err error
	if _, err := s.storage.GetUserByLogin(user.Username); err == nil {
		return fmt.Errorf("service: create user: %w: username already used", ErrUserExist)
	}
	if _, err := s.storage.GetUserByEmail(user.Email); err == nil {
		return fmt.Errorf("service: create user: %w: email already used", ErrUserExist)
	}
	if err := validUser(user); err != nil {
		return fmt.Errorf("service: create user: %w", err)
	}

	user.Password, err = generateHashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("service: create user: %w", err)
	}
	return s.storage.CreateUser(user)
}

func (s *AuthService) GenerateSessionToken(username string) (string, error) {
	user, err := s.storage.GetUserByLogin(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("service: generate session token: %w", ErrUserNotFound)
		}
		return "", fmt.Errorf("service: generate session token: %w", err)
	}

	token := uuid.NewString()

	if err := s.storage.SaveSessionToken(user.Username, token); err != nil {
		return "", fmt.Errorf("service: generate session token: %w", err)
	}
	return token, nil
}

func (s *AuthService) ParseSessionToken(token string) (models.User, error) {
	user, err := s.storage.GetUserByToken(token)
	if err != nil {
		return user, fmt.Errorf("service: patse session token: %w", err)
	}
	return user, nil
}

func (s *AuthService) DeleteSessionToken(token string) error {
	err := s.storage.DeleteSessionToken(token)
	if err != nil {
		return fmt.Errorf("service: delete session token: %w", err)
	}
	return nil
}

func generateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func compareHashAndPassword(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return ErrUserNotFound
	}
	return nil
}

func validUser(user models.User) error {
	for _, char := range user.Username {
		if char <= 32 || char >= 127 {
			return ErrInvalidUserName
		}
	}
	validEmail, err := regexp.MatchString(`[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`, user.Email)
	if err != nil {
		return err
	}
	if !validEmail {
		return ErrInvalidEmail
	}
	if len(user.Username) <= 4 || len(user.Username) >= 36 {
		return ErrInvalidUserName
	}

	return nil
}
func (s *AuthService) CheckToken(token string) error {
	if token == "" {
		return errors.New("no tokens")
	}
	return s.storage.CheckToken(token)
}

func (s *AuthService) GetUserByToken(token string) (models.User, error) {
	return s.storage.GetUserByToken(token)
}
