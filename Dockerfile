# ---------------------
# build envoy
FROM golang:1.20-bullseye as envoy-backend

## set env
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

## build
WORKDIR /opt/envoy/envoy
RUN cd  /opt/envoy/envoy
RUN ls -alh

COPY ./ ./

RUN cat ./Makefile

RUN make build-http-server

RUN ls -alh ./bin/envoy


# -------------------
# build runner images
FROM alpine:latest as runner

WORKDIR /opt/envoy/envoy/bin/

## copy backend bin
COPY --from=envoy-backend /opt/envoy/envoy/bin/envoy /opt/envoy/envoy/bin/


RUN ls -alh /opt/envoy/envoy/bin/



# run
EXPOSE 8001
CMD ["/bin/sh", "-c", "/opt/envoy/envoy/bin/envoy"]
