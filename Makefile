download:
	go mod download

install-tools: download
	@echo Installing tools from tools.go
	cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

format-all:
	go fmt ./...

lint: format-all
	golangci-lint run ./...

doc:
	godoc -http=:6060

test:
	go test -v ./...

benchmark:
	go test -bench . ./...

test-cover:
	go test -v -coverprofile cover.out ./...
	go tool cover -func cover.out
	go tool cover -html cover.out

vet:
	go vet ./...

build: download install-tools lint test vet
	go build ./...