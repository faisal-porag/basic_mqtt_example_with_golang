docker_up:
	docker-compose up -d


run_server:
	go run main.go

run_client:
	go run client/main.go

