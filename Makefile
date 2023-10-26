run-job:
	@set -a && source ./.env && set +a && go run cmd/job/main.go
run-server:
	@set -a && source ./.env && set +a && go run cmd/server/main.go
migrate:
	@set -a && source ./.env && set +a && go run cmd/migrations/main.go
fmt:
	go fmt ./...
lint:
	golangci-lint run
test:
	go test -v ./...
setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2 && go mod tidy
buf-push:
	buf push && buf mod update && buf build && go get buf.build/gen/go/capybara/my-assemble/connectrpc/go@latest && go get buf.build/gen/go/capybara/my-assemble/protocolbuffers/go@latest
buf-lint:
	buf lint
buf-fmt:
	buf format -w