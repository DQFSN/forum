FROM golang:1.52-stretch
WORKDIR /go/src/github.com/DQFSN/blog
COPY . .
RUN go build -o /blog/server .
CMD ["/blog/server"]