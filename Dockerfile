FROM golang

ADD . /go/src/github.com/adorsys/gitlab-reporting

WORKDIR /go/src/github.com/adorsys/gitlab-reporting

RUN go install

EXPOSE 9090

ENTRYPOINT /go/bin/gitlab-reporting
