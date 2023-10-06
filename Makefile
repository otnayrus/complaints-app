docker-up:
	@docker-compose up --build -d

run-rest:
	@make -s docker-up
	@go run -race -v app/main.go

run-client:
	@npm start