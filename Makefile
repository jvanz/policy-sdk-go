# Run tests
test: fmt vet
	go test ./... -coverprofile cover.out

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

lint:
	golangci-lint run

deps:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint

generate-mocks:
	mockery
