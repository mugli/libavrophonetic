download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

format-all:
	@go fmt ./...

lint: install-tools format-all
	@echo golangci-lint run
	@golangci-lint run ./...