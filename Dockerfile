FROM golang
WORKDIR /go/src/github.com/DQFSN/forum
COPY . .
RUN go build -o /blog/server .
CMD ["/blog/server"]