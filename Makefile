.PHONY: build run migrate

build:
	docker compose build

run:
	docker compose up -d

down:
	docker compose down 

migrate:
	docker compose exec ./todo-app migrate