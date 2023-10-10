docker-up:
	@docker-compose up --build -d

docker-down:
	@docker-compose down

run-server:
	@go mod vendor
	@make -s docker-up
	@go run -race -v app/main.go

run-client:
	@yarn install
	@yarn start

install:
	@go mod vendor
	@yarn install
