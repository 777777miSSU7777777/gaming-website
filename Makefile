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
	go build

.PHONY: run
run: 
	sudo docker-compose up --build

.PHONY: docker
docker: ;