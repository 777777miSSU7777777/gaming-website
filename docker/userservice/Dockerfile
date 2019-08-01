# Build stage
FROM golang:1.12-alpine3.10 AS builder

ENV GO111MODULE=on

ADD . /src

RUN apk add build-base \
    && apk update \
    && apk upgrade \
    && apk add git

RUN cd /src && go build cmd/userservice/user_service.go

# Final stage 
FROM alpine:latest AS runtime

EXPOSE 8080

WORKDIR /app

COPY --from=builder /src/user_service /app/

CMD ["sh", "-c", "./user_service -connection_string=$CONNECTION_STRING"]