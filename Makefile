all: install check

install:
	@echo "+ $@"
	@go get -t ./...
	@go get github.com/golang/lint/golint

check:
	@echo "+ $@"
	@go vet ./...
	@test -z "$(golint ./... | tee /dev/stderr)"
	@test -z "$(gofmt -s -l . | tee /dev/stderr)"
	@go test ./...

