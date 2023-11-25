build:
	go build -o application .
run:
	docker compose down && docker compose build && docker compose up
	

migrate-up:
	docker exec -ti queryMaster_app ./application migrate up

migrate-down:
	docker exec -ti queryMaster_app ./application migrate down

docker-exec-psql:
	docker exec -ti querymaster_postgres psql -U username querymaster
