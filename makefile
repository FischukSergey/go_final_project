# устанавливаем переменные среды окружения
port:=:7540
envRunAddr:=TODO_PORT=$(port)
envDatabaseDSN:=TODO_DBFILE=./storage/scheduler.db
envPassword:=TODO_PASSWORD=12345
server:
				@echo "Running server"
				$(envRunAddr) $(envDatabaseDSN) $(envPassword) go run ./cmd/*.go
.PHONY: server

TestDBserver:
				@echo "Running default server "
				go build -o ./cmd/final_project ./cmd/*.go
				$(envRunAddr) $(envDatabaseDSN) $(envPassword) ./cmd/final_project
.PHONY: TestDBserver

test:
				@echo "Running unit tests"
				go test -race -count=1 -cover ./...
.PHONY: test

autotest:
				@echo "Running unit tests"
				go test -count=1 -run ^TestApp$  ./tests
				go test -count=1 -run ^TestDB$  ./tests
				go test -count=1 -run ^TestNextDate$  ./tests
				go test -count=1 -run ^TestAddTask$  ./tests
				go test -count=1 -run ^TestTasks$  ./tests
				go test -count=1 -run ^TestTask$  ./tests
				go test -count=1 -run ^TestEditTask$  ./tests
				go test -count=1 -run ^TestDone$  ./tests
				go test -count=1 -run ^TestDelTask$  ./tests
.PHONY: autotest

statictest:
				@echo "Running static tests"
				go vet ./...
				go test -v ./internal/lib
.PHONY:statictest

testcover:
				@echo "Running unit tests into file"
				go test -coverprofile=coverage.out ./...
				go tool cover -func=coverage.out
.PHONY: testcover

build:
				go build -o ./cmd/final_project ./cmd/*.go
.PHONY: build

dockerbuild:
	@echo "Running create docker image"
	docker build -t diplom:v0.0.1 .
.PHONY: dockerbuild

dockerrun:
	@echo "Running create docker image"
	docker run diplom:v0.0.1
.PHONY: dockerrun


# curl -v -X GET 'http://localhost:8080/map'
# curl -v -d "http://yandex.ru" -X POST 'http://localhost:8080/'
# curl -v -d '{"url": "https://codewars.com"}' -H "Content-Type: application/json" POST 'http://localhost:8080/api/shorten'
# curl -v -X GET 'http://localhost:8080/map' -H "Accept-Encoding: gzip"
# /Users/sergeymac/dev/urlshortener/shortenertestbeta-darwin-arm64 -test.v -test.run=^TestIteration9$ -binary-path=cmd/shortener/shortener -file-storage-path=tmp/short-url-db.json -source-path=tmp/short-url-db.json -database-dsn=urlshortdb
# pg_ctl -D /usr/local/pgsql/data stop/start
# go build -o shortener *.go