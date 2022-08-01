APP_BUILD = build/apiserver

.DEFAULT_GOAL = build

.PHONY: build
build:
	clear; go build -o ${APP_BUILD} -v ./cmd/app/main.go

.PHONY: start
start: build
	${APP_BUILD}

.PHONY: migrate-dev
migrate-dev:
	migrate -path migrations -database "postgres://math:math@localhost/camp_dev?sslmode=disable" ${type}