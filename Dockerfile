FROM golang

RUN mkdir -p /go/src/github.com/johnmcdnl/flashcards
WORKDIR /go/src/github.com/johnmcdnl/flashcards
COPY . .
RUN go build -o flashcards ./cmd/main/main.go

FROM debian:stretch-slim
COPY --from=0  /go/src/github.com/johnmcdnl/flashcards/flashcards ./flashcards
COPY --from=0  /go/src/github.com/johnmcdnl/flashcards/data ./data
CMD ["./flashcards"]
