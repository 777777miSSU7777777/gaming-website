# Build stage
FROM golang:1.12 AS builder

ENV GO111MODULE=on

ADD . /src

RUN cd /src && go build main.go

# Final stage 
FROM ubuntu:18.04 AS runtime

EXPOSE 8080

WORKDIR /app

COPY --from=builder /src/main /app/

CMD ["sh", "-c", "./main -connection_string=$CONNECTION_STRING"]