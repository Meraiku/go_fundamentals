build:
	@go build -o ./.bin/out ./cmd/out/main.go

run:build
	@./.bin/out
