# Nick Turner
# Porch.com

FROM golang
MAINTAINER Nick Turner

ADD . /go/src/k8s-test-service

RUN go install /go/src/k8s-test-service

ENTRYPOINT /go/bin/k8s-test-service

EXPOSE 11011
