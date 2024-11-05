build:
	@go build -o ./.bin/out ./cmd/out/main.go

run:build
	@./.bin/out

tests:
	@go test ./... -v -race 

cover:
	@go test ./... -v -race -cover -coverprofile=coverage.out
	@go tool cover -html=coverage.out
