.PHONY: fmt lint run

fmt:
	gofmt -s -l -w .

lint: fmt
	golangci-lint run

run:
	go run *.go