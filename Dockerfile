FROM golang:alpine AS builder
WORKDIR ./src/github.com/johnmcdnl/flashcards
ADD . .
RUN go build -o flashcards ./cmd/main/main.go

FROM alpine
WORKDIR /flashcards
COPY --from=builder /go/src/github.com/johnmcdnl/flashcards/flashcards .
COPY ./data ./data
ENTRYPOINT ./flashcards
