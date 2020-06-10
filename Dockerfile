FROM golang:1.13-alpine AS builder

RUN apk update && \
    apk upgrade && \
    apk add --no-cache curl make git \
    && curl -fLo /usr/local/bin/air https://git.io/linux_air  \
    && chmod +x /usr/local/bin/air \
    && mkdir -p /app \
    && export GOPRIVATE="github.com/fingerpeople"
    
WORKDIR /app
COPY    . .

EXPOSE 3000