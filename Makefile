run:
	go run cmd/main.go

up:
	docker-compose -f docker-compose.yaml up --build

down:
	docker-compose -f docker-compose.yaml down

.PHONY: run up down