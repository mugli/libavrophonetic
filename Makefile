download:
	go mod download
	go mod tidy

install-build-tools: download
	@echo Installing tools from buildtools.go
	cat internal/buildtools/buildtools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
	go mod tidy

format-all:
	go fmt ./...
	gci -w .

lint: format-all
	golangci-lint run ./...

doc:
	godoc -http=:6060

test:
	go test ./...

benchmark:
	go test -bench . ./...

test-cover:
	go test -coverprofile cover.out ./...
	go tool cover -func cover.out
	go tool cover -html cover.out

vet:
	go vet ./...

generate-data:
	go run ./internal/cmd/generate_data/ -data-directory=./data
	go run ./internal/cmd/generate_data/ -data-directory=./databasedconv/testdata

build: download install-build-tools lint test vet
	go build ./...