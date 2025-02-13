# ---------------------
# build envoy-builder
FROM golang:1.20-bullseye as builder-for-backend

## set env
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

## build
WORKDIR /opt/envoy/envoy-builder
RUN cd  /opt/envoy/envoy-builder
RUN ls -alh

COPY ./ ./

RUN cat ./Makefile

RUN make build-http-server

RUN ls -alh ./bin/envoy-builder


# -------------------
# build runner images
FROM alpine:latest as runner

WORKDIR /opt/envoy/envoy-builder/bin/

## copy backend bin
COPY --from=builder-for-backend /opt/envoy/envoy-builder/bin/envoy-builder /opt/envoy/envoy-builder/bin/


RUN ls -alh /opt/envoy/envoy-builder/bin/



# run
EXPOSE 8001
CMD ["/bin/sh", "-c", "/opt/envoy/envoy-builder/bin/envoy-builder"]
