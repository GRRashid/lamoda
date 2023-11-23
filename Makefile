build:
	sudo docker-compose build lamoda

up: build
	sudo docker-compose up lamoda

migrate:
	goose -dir ./migrations/ postgres "postgresql://postgres:qwerty@localhost:5436/postgres" up