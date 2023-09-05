run-producer:
	go run ./cmd/producer/main.go

run-consumer:
	go run ./cmd/consumer/main.go

docker:
	docker-compose up -d