FROM golang

ADD proto/calc.pb.go /go/src/github.com/tassygan/end/proto/
ADD server/main.go /go/src/github.com/tassygan/end/server/

RUN go get -u golang.org/x/net/context
RUN go get -u google.golang.org/grpc
RUN go get -u google.golang.org/grpc/reflection

RUN go install github.com/tassygan/end/server

ENTRYPOINT ["/go/bin/server"]

EXPOSE 50051
