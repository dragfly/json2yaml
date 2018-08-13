FROM golang:alpine

RUN apk --no-cache add git

ADD . "$GOPATH/src/json2yaml"

RUN go get github.com/ghodss/yaml
RUN go install json2yaml

FROM alpine:latest

WORKDIR /bin/

COPY --from=0 /go/bin/json2yaml .

WORKDIR /src

ENTRYPOINT ["/bin/json2yaml"]
