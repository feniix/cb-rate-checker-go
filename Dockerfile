FROM golang:1.9.2-alpine3.7 as stage
RUN apk add --no-cache git
RUN \
  wget -q https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 -O /usr/local/bin/dep && \
  chmod 0755 /usr/local/bin/dep


WORKDIR /go/src/github.com/feniix/cb-rate-checker-go

COPY ./Gopkg.* ./
COPY ./*.go ./

RUN \
  dep ensure && \
  dep status && \
  CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .

FROM alpine:3.7

RUN apk add --no-cache ca-certificates

COPY --from=stage /go/src/github.com/feniix/cb-rate-checker-go/cb-rate-checker-go /rate-checker

RUN chmod a+x /rate-checker

ENTRYPOINT [ "/rate-checker" ]
