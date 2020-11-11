FROM golang:1.11

WORKDIR $GOPATH/src/jarvan
COPY . $GOPATH/src/jarvan

RUN go build .

EXPOSE 8000
ENTRYPOINT ["./jarvan"]