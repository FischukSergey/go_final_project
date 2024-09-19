FROM golang:1.22
WORKDIR /app
#копируем все файлы, включая DB
COPY . .

#грузим зависимости
RUN go mod download

#монтируем приложение
RUN go build -o /diplom /app/cmd/*.go

#определяем переменные среды окружения
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV TODO_PORT=:7540
ENV TODO_DBFILE=/app/storage/scheduler.db
ENV LOG_LEVEL=local

# Запускаем приложение
CMD ["/diplom"]
