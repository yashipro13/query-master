build:
	go build -o application .
run:
	docker compose down && docker compose build && docker compose up