.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	goimports -w ./

.PHONY: test
test: ;

.PHONY: build
build:
	go build cmd/server/server.go

.PHONY: run
run: ;

.PHONY: docker
docker: ;