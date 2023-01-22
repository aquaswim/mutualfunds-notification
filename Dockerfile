FROM golang:alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o aya cmd/aya.go

FROM alpine
RUN apk add --no-cache tzdata
COPY --from=builder /app/aya /usr/local/bin/aya


ENTRYPOINT ["aya"]