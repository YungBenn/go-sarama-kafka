run-producer:
	go run ./producer/main.go

run-consumer:
	go run ./consumer/main.go

docker:
	docker-compose up -d