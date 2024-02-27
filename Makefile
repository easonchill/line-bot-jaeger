include .env
EXECUTABLE := main
GOFILES := $(shell find . -type f -name "*.go")

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

build: $(EXECUTABLE)
$(EXECUTABLE): $(GOFILES)
	go build -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o line-bot-jaeger ./main.go

db-up:
	DB_USERNAME=${DB_USERNAME}
	DB_PASSWORD=${DB_PASSWORD}
	DB_PORT=${DB_PORT}
	docker-compose up -d
db-down:
	docker-compose down -v
test:
	go test -v -cover -coverprofile coverage.txt ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1
