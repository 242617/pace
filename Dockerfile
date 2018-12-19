FROM golang:1.9

ENV BINARY pace

ARG APPLICATION
ARG PROJECT
ARG SERVER_ADDRESS

ENV APPLICATION ${APPLICATION}
ENV PROJECT ${PROJECT}
ENV SERVER_ADDRESS ${SERVER_ADDRESS}

WORKDIR $GOPATH
COPY . src/github.com/242617/pace

RUN go get ./...
RUN go build \
    -ldflags " \
        -X ${PROJECT}/version.Application=${APPLICATION} \
        -X ${PROJECT}/config.ServerAddress=${SERVER_ADDRESS} \
    " \
    -o ${BINARY} \
    ${PROJECT}/cmd/${APPLICATION}

CMD ["sh", "-c", "$BINARY"]