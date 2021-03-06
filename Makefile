download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

format-all:
	@echo Formatting files
	@go fmt ./...

lint: format-all
	@echo golangci-lint run
	@golangci-lint run ./...

doc:
	@echo godoc -http=:6060
	@godoc -http=:6060

test:
	@echo go test ./...
	@go test ./...

vet:
	@echo go vet ./...
	@go vet ./...

build: download install-tools lint test vet
	@echo go build ./...
	@go build ./...