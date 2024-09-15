package models

// ограничения на количество задач в ответе
const (
	LimitTasks = 10 //ограничение на количество задач в ответе
)

// Task структура задачи из запроса
type Task struct {
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

// SaveTask структура для сохранения задачи
type SaveTask struct {
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// SearchTask структура для поиска задачи
type SearchTask struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// SearchTasksResponse структура для ответа на запрос поиска задачи
type SearchTasksResponse struct {
	Tasks []SearchTask `json:"tasks"`
	Error string       `json:"error,omitempty"`
}

// ErrorResponse структура для ответа на запрос с ошибкой
type ErrorResponse struct {
	Error string `json:"error"`
}	