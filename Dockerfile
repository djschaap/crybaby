FROM docker.io/golang:1.21-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum *.go ./

RUN \
  go build -o crybaby crybaby.go

FROM scratch

USER 1000

COPY --from=builder /app/crybaby /crybaby

CMD ["/crybaby"]
