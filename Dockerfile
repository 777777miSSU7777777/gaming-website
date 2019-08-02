# Build stage
FROM golang:1.12 AS builder

ENV GO111MODULE=on

ADD . /src

RUN cd /src && go build cmd/userservice/user_service.go

# Final stage 
FROM ubuntu:18.04 AS runtime

EXPOSE 8080

WORKDIR /app

COPY --from=builder /src/user_service /app/

CMD ["sh", "-c", "./user_service -connection_string=$CONNECTION_STRING"]