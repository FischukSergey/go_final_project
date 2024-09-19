package register

import (
	"log/slog"
	"net/http"

	"github.com/FischukSergey/go_final_project/internal/lib/jwtoken"
	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

// Register функция для аутентификации пользователя
func Register(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//структура для ответа
		type RegisterToken struct {
			Token string `json:"token"`
		}
		//структура для получения пароля
		type Register struct {
			Password string `json:"password"`
		}
		password := Register{}
		err := render.DecodeJSON(r.Body, &password) //получаем пароль из тела запроса
		if err != nil {
			log.Error("ошибка при декодировании пароля", logger.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
			return
		}

		pass := models.Pass //получаем пароль из переменной окружения
		if pass == "" {     //если пароль не установлен, то возвращаем ошибку
			log.Error("пароль не установлен")
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, models.ErrorResponse{Error: "пароль не установлен"})
			return
		}

		if password.Password != pass { //если пароль не совпадает, то возвращаем ошибку
			log.Error("неверный пароль")
			w.WriteHeader(http.StatusUnauthorized)
			render.JSON(w, r, models.ErrorResponse{Error: "неверный пароль"})
			return
		}

		//хэшируем пароль
		passHash, err := bcrypt.GenerateFromPassword([]byte(password.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error("ошибка при генерации хэша пароля", logger.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
			return
		}

		//создаем пользователя
		user := models.User{
			ID:                1,                //TODO: изменить на ID пользователя
			Login:             "admin",          //TODO: изменить на логин пользователя
			EncryptedPassword: string(passHash), //пишем хэш пароля
		}

		token, err := jwtoken.NewToken(user) //генерируем токен
		if err != nil {                      //если ошибка при генерации токена, то возвращаем ошибку
			log.Error("ошибка при генерации токена", logger.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
			return
		}

		//если все успешно, то пишем токен в ответе
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, RegisterToken{Token: token})
	}
}
