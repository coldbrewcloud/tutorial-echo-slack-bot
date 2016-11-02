FROM golang:1.7.3

COPY . /go/src/github.com/coldbrewcloud/tutorial-echo-slack-bot
RUN cd /go/src/github.com/coldbrewcloud/tutorial-echo-slack-bot && \
    GOPATH=/go go get -d -v && \
    GOPATH=/go go install -v

CMD ["/go/bin/tutorial-echo-slack-bot"]