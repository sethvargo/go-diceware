test:
	@go test -parallel=20 ./...
.PHONY: test

test-race:
	@go test -parallel=20 -race ./...
.PHONY: test-race
