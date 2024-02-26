include .env

db-up:
	DB_USERNAME=${DB_USERNAME}
	DB_PASSWORD=${DB_PASSWORD}
	DB_PORT=${DB_PORT}
	docker-compose up -d
db-down:
	docker-compose down -v