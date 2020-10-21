FROM golang
WORKDIR /go/src/github.com/DQFSN/forum
COPY . .
RUN apt update
RUN apt install zip -y
RUN mkdir -p /go/pkg/mod/github.com/micro/go-plugins/
RUN cd /go/pkg/mod/github.com/micro/go-plugins/ && wget https://github.com/asim/go-plugins/archive/wrapper/validator/v2.9.1.zip
RUN cd /go/pkg/mod/github.com/micro/go-plugins/ && unzip v2.9.1.zip
RUN cd /go/pkg/mod/github.com/micro/go-plugins/ && mv go-plugins-wrapper-validator-v2.9.1/ v2@v2.0.0
RUN go mod tidy

RUN go run server/start.go