FROM golang

RUN mkdir -p /go/src/github.com/johnmcdnl/flashcards
WORKDIR /go/src/github.com/johnmcdnl/flashcards
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=0  /go/src/github.com/johnmcdnl/flashcards/main /main
CMD ["/main"]
