FROM alpine:latest

ENV TIMEZONE=Europe/Istanbul
EXPOSE 9090 9092 6060

# setup timezone
RUN apk update && apk upgrade && \
    apk add --no-cache ca-certificates && \
    apk add -U --no-cache tzdata && \
    cp /usr/share/zoneinfo/$TIMEZONE /etc/localtime && echo $TIMEZONE > /etc/timezone && \
    rm -rf /var/cache/apk/*

COPY ./config /app/config
COPY url-shortener /usr/local/bin/url-shortener
WORKDIR /app

ENTRYPOINT ["/usr/local/bin/url-shortener"]
