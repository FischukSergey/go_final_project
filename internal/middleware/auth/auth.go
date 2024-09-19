package auth

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/FischukSergey/go_final_project/internal/lib/jwtoken"
	"github.com/FischukSergey/go_final_project/internal/logger"
	"github.com/FischukSergey/go_final_project/internal/models"
	"github.com/go-chi/render"
)

// CtxKey тип для ключей контекста
type ctxKey int

// CtxKeyUser ключ для пользователя
const (
	CtxKeyUser ctxKey = iota + 1
)

// AuthToken middleware проверка на авторизацию
func AuthToken(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		log.Debug("middleware авторизация")

		Authorize := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			pass := models.Pass

			if len(pass) == 0 { //если пароль не установлен, то аутентификацию не проводим
				log.Debug("пароль не установлен")
				next.ServeHTTP(w, r)
			} else {
				log.Debug("пароль установлен, проводим аутентификацию")
				token, err := r.Cookie("token") //получаем токен из куки
				if err != nil {                 //если токена нет, то возвращаем ошибку
					log.Error("токен не найден", logger.Err(err))
					w.WriteHeader(http.StatusUnauthorized)
					render.JSON(w, r, models.ErrorResponse{Error: "необходимо авторизоваться"})
					return
				}

				userID, err := jwtoken.GetJWTokenUserID(token.Value) //проверяем токен на валидность и получаем ID пользователя
				if err != nil {                                      //токен есть, но не валиден
					log.Error("токен не валиден", logger.Err(err))
					w.WriteHeader(http.StatusUnauthorized)
					render.JSON(w, r, models.ErrorResponse{Error: err.Error()})
					return
				}

				//если все успешно - пишем в контекст ID пользователя для последующего использования в хендлерах
				log.Info("токен валиден", slog.String("user ID:", strconv.Itoa(userID)))
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), CtxKeyUser, userID)))
			}
		}
		return http.HandlerFunc(Authorize)
	}
}
