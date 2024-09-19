package jwtoken

import (
	"errors"
	"fmt"
	
	"github.com/FischukSergey/go_final_project/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	secretkey = "very-secret-key"
)

// NewToken генерируем JWToken
func NewToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	if user.ID > 0 && user.Login != "" {
		claims["uid"] = user.ID
		claims["login"] = user.Login
		claims["encrypted_password"] = user.EncryptedPassword
	} else {
		return "", errors.New("невозможно создать токен, неверный ID или логин")
	}

	tokenString, err := token.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// проверка валидности токена
func GetJWTokenUserID(tokenString string) (int, error) {

	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretkey), nil
		})
	if err != nil {
		return -1, fmt.Errorf("ошибка при парсинге токена: %w", err)
	}

	if !token.Valid {
		return -1, fmt.Errorf("токен невалидный")
	}

	userID := claims["uid"].(float64)
	encryptedPassword := claims["encrypted_password"].(string)
	//проверяем не изменился ли пароль с момента создания токена
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(models.Pass)); err != nil {
		return -1, fmt.Errorf("пароль изменился	: %w", err)
	}

	return int(userID), nil
}
