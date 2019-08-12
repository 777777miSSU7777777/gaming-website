.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	goimports -w ./

.PHONY: test
test: 
	sudo docker-compose up --build -d
	go test test/main_test.go
	sudo docker-compose down

.PHONY: build
build:
	go build

.PHONY: run
run: 
	sudo docker-compose up --build

.PHONY: docker
docker: ;