FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY powers /

EXPOSE 8013

ENTRYPOINT ["/powers"]

