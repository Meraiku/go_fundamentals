build:
	@go build -o ./.bin/out ./cmd/out/main.go

run:build
	@./.bin/out

docker:build
	@docker compose up -d --build

stop:
	@docker compose stop

unit:
	@go test $(shell go list ./... | grep -v /tests)  -v -race 

integration:
	@go test ./tests/... -v -race

cover:
	@go test $(shell go list ./... | grep -v /tests) -v -race -cover -coverprofile=coverage.out
	@go tool cover -html=coverage.out

mocks:
	@go get github.com/golang/mock/mockgen
	@mockgen -source=internal/client/client.go -destination=internal/client/mocks/mock_client.go
