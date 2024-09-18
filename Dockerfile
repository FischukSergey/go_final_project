FROM golang:1.22
WORKDIR /app
#копируем все файлы, включая DB
COPY . .
RUN ls -la
#грузим зависимости
RUN go mod tidy

#монтируем приложение
RUN CGO_ENABLED=1 GOOS=linux go build -o /diplom /app/cmd/*.go

# Запускаем приложение с использованием переменных среды
#CMD ["/diplom", "TODO_PORT=127.0.0.1:7540", "TODO_DBFILE=/app/storage/scheduler.db", "LOG_LEVEL=local"]
CMD ["/diplom"]
