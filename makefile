# устанавливаем переменные среды окружения
port:=:7540
envRunAddr:=TODO_PORT=$(port)
#envDatabaseDSN:=DATABASE_DSN="user=postgres password=postgres host=localhost port=5432 dbname=urlshortdb sslmode=disable"

server:
				@echo "Running server"
				$(envRunAddr)  go run ./cmd/*.go
.PHONY: server

db:
				@echo "Running server"
				$(envRunAddr) $(envBaseURL) $(envDatabaseDSN) go run ./cmd/shortener/main.go
.PHONY: server

defaultserver:
				@echo "Running default server "
				go run ./cmd/main.go

test:
				@echo "Running unit tests"
				go test -race -count=1 -cover ./...
.PHONY: test

autotest:
				@echo "Runing autotest"
				go test -count=1 -run ^TestApp$ ./tests
.PHONY: autotest

testcover:
				@echo "Running unit tests into file"
				go test -coverprofile=coverage.out ./...
				go tool cover -func=coverage.out
.PHONY: testcover

build:
				go build -o ./cmd/final_project ./cmd/*.go
.PHONY: build


# curl -v -X GET 'http://localhost:8080/map'
# curl -v -d "http://yandex.ru" -X POST 'http://localhost:8080/'
# curl -v -d '{"url": "https://codewars.com"}' -H "Content-Type: application/json" POST 'http://localhost:8080/api/shorten'
# curl -v -X GET 'http://localhost:8080/map' -H "Accept-Encoding: gzip"
# /Users/sergeymac/dev/urlshortener/shortenertestbeta-darwin-arm64 -test.v -test.run=^TestIteration9$ -binary-path=cmd/shortener/shortener -file-storage-path=tmp/short-url-db.json -source-path=tmp/short-url-db.json -database-dsn=urlshortdb
# pg_ctl -D /usr/local/pgsql/data stop/start
# go build -o shortener *.go