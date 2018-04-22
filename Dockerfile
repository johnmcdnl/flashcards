FROM golang:alpine AS builder
WORKDIR ./src/github.com/johnmcdnl/flashcards
ADD . .
RUN go build -o flashcards-api ./cmd/api/api.go

FROM alpine
WORKDIR /flashcards
COPY --from=builder /go/src/github.com/johnmcdnl/flashcards/flashcards-api .
COPY ./data ./data

EXPOSE 5800

ENTRYPOINT ./flashcards-api