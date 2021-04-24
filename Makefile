download:
	go mod download
	go mod tidy

install-build-tools: download
	@echo Installing tools from buildtools.go
	cat buildtools/buildtools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
	go mod tidy

format-all:
	go fmt ./...
	gci -w .

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

generate-data:
	go run ./tools/cmd/generate_data/ -data-directory=./data

build: download install-build-tools lint test vet
	go build ./...