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
	go build cmd/userservice/user_service.go

.PHONY: run
run: 
	sudo docker-compose build
	sudo docker-compose up

.PHONY: docker
docker: ;