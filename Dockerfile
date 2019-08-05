# Build stage
FROM golang:1.12 AS builder

ENV GO111MODULE=on

ADD . /src

RUN cd /src && go build -o service main.go

# Final stage 
FROM ubuntu:18.04 AS runtime

EXPOSE 8080

WORKDIR /app

COPY --from=builder /src/service /app/

CMD ["sh", "-c", "./service -connection_string=$CONNECTION_STRING"]