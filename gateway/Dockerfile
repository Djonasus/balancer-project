# Этап, на котором выполняется сборка приложения
FROM golang:1.23-alpine as builder
WORKDIR /build
COPY go.mod .
# COPY config.json /build/config.json
RUN go mod download
COPY . .
RUN go build -o /main .
# Финальный этап, копируем собранное приложение
FROM alpine:3
COPY --from=builder main /bin/main
# COPY --from=builder config.json /bin/config.json
ENTRYPOINT ["/bin/main"]