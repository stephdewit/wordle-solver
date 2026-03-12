FROM golang:1.25-alpine AS builder
WORKDIR /src
COPY go.mod .
COPY solver/ solver/
COPY api/ api/
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /wordle-solver-api ./api

FROM alpine:3.21
RUN apk add --no-cache words
COPY --from=builder /wordle-solver-api /wordle-solver-api
EXPOSE 8080
ENTRYPOINT ["/wordle-solver-api"]
