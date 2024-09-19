package nextdate

import (
	"log/slog"
	"net/http"
	"time"

	repeatrule "github.com/FischukSergey/go_final_project/internal/lib"
	"github.com/FischukSergey/go_final_project/internal/logger"
)

// NextDate - обработчик для получения следующей даты задачи
func NextDate(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Next date api started")
		defer log.Debug("Next date api finished")
		
		
		date := r.URL.Query().Get("date")
		repeat := r.URL.Query().Get("repeat")
		now := r.URL.Query().Get("now")
		if now == "" || date == "" || repeat == "" {
			log.Error("Invalid request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		
		nowDate, err := time.Parse("20060102", now) //парсим дату в формате YYYYMMDD
		if err != nil {
			log.Error("Invalid date now", "error", err)
			http.Error(w, "Invalid date now", http.StatusBadRequest)
			return
		}
		
		nextDate, err := repeatrule.NextDate(nowDate, date, repeat)
		if err != nil {
			log.Error("Invalid date", "error", err)
			http.Error(w, "Invalid date", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(nextDate))
		if err != nil {
			log.Error("Ошибка записи ответа", logger.Err(err))
			http.Error(w, "Ошибка записи ответа", http.StatusInternalServerError)
			return
		}
	}
}