.PHONY: test coverage

# Run tests for the cmd/web folder
test:
	go test ./cmd/web/...

# Generate HTML coverage report for the cmd/web folder
coverage: test
	go test -coverprofile=coverage.out ./cmd/web/...
	go tool cover -html=coverage.out -o coverage.html
