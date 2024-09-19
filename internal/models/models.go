package models

// ограничения на количество задач в ответе
const (
	LimitTasks = 10 //ограничение на количество задач в ответе
)

// Pass пароль для доступа к API
var Pass string

// Task структура задачи из запроса
type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// TaskResponse структура ответа на запрос
type TaskResponse struct {
	ID    string `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

// SearchTasksResponse структура для ответа на запрос поиска задачи
type SearchTasksResponse struct {
	Tasks []Task `json:"tasks"`
	Error string       `json:"error,omitempty"`
}

// ErrorResponse структура для ответа на запрос с ошибкой
type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

// User структура для пользователя
type User struct {
	ID                int    `json:"id"`
	Login             string `json:"login"`
	Password          string `json:"password"`
	EncryptedPassword string `json:"encrypted_password"`
}
