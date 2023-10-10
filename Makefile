.PHONY: run
run:
	@set -a && source ./.env && set +a && go run cmd/job/main.go
.PHONY: fmt
fmt:
	go fmt ./...
lint:
	golangci-lint run
test:
	go test -v ./...