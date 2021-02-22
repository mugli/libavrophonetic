download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

format-all:
	@echo Formatting files
	@gofumpt -w .

lint: format-all
	@echo golangci-lint run
	@golangci-lint run ./...