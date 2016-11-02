FROM alpine:3.4

RUN apk --update add ca-certificates

COPY bot /

EXPOSE 8888

CMD ["/bot"]