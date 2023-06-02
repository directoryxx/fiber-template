FROM golang:1.20-alpine as builder
LABEL authors="directoryx"

RUN mkdir /app

RUN apk add build-base librdkafka-dev pkgconf

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=1 go build -tags musl -o ./build/app ./cmd/app

RUN CGO_ENABLED=1 go build -tags musl -o ./build/app-cli ./cmd/cli

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN chmod +x /app/build/app

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/build/app /app
COPY --from=builder /app/build/app-cli /app
COPY --from=builder /app/pkg /app/pkg
COPY --from=builder /go/bin/migrate /bin/migrate

# COPY ./.env /.env

CMD [ "/app/app" ]
EXPOSE 8000