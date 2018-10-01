FROM golang:1.11.0

WORKDIR $GOPATH/src/github.com/berto/excel-processor
COPY ./* ./

ENTRYPOINT ["/bin/sh"]
