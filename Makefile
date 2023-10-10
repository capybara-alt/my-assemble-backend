run:
	@set -a && source ./.env && set +a && go run cmd/job/main.go
fmt:
	go fmt ./...
lint:
	golangci-lint run
test:
	go test -v ./...
setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2 && go mod tidy