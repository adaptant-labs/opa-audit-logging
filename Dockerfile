FROM golang:latest as builder

ENV GO111MODULE=on

ARG AUDIT_LOGGING_SERVICE_PORT
ARG AUDIT_LOGGING_SERVICE_MONGO_URI

WORKDIR /go/src
ADD . /go/src

RUN go mod download
RUN go build -ldflags "-linkmode external -extldflags -static" -a -o /go/bin/app

FROM scratch
COPY --from=builder /go/bin/app /

EXPOSE 4123
ENTRYPOINT ["/app" ]
