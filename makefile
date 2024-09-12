# устанавливаем переменные среды окружения
port:=:7540
envRunAddr:=TODO_PORT=$(port)
envDatabaseDSN:=TODO_DBFILE=./storage/scheduler.db

server:
				@echo "Running server"
				$(envRunAddr) $(envDatabaseDSN) go run ./cmd/*.go
.PHONY: server

TestDBserver:
				@echo "Running default server "
				go build -o ./cmd/final_project ./cmd/*.go
				$(envRunAddr) $(envDatabaseDSN) ./cmd/final_project
.PHONY: TestDBserver

test:
				@echo "Running unit tests"
				go test -race -count=1 -cover ./...
.PHONY: test

autotest:
				@echo "Running unit tests"
				go test -v ./lib/.
				go test -count=1 -run ^TestApp$  ./tests
				go test -count=1 -run ^TestDB$  ./tests
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