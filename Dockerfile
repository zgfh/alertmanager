FROM harbor.geniusafc.com/docker.io/golang:1.12 as builder
ENV GO111MODULE=on
ENV GOPROXY http://nexus.geniusafc.com/repository/go/
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO15VENDOREXPERIMENT 1

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN cd cmd/alertmanager; go build -o /bin/alertmanager *.go; cd ../amtool ; go build -o /bin/amtool *.go

FROM harbor.geniusafc.com/infra/alpine:3.10
WORKDIR /app
COPY --from=builder /bin/app  /bin/app
EXPOSE 9093

CMD ["app"]

FROM quay.io/prometheus/busybox-amd64-linux:latest
LABEL maintainer="The Prometheus Authors <prometheus-developers@googlegroups.com>"

COPY --from=builder /bin/amtool  /bin/amtool
COPY --from=builder /bin/alertmanager  /bin/alertmanager
COPY examples/ha/alertmanager.yml      /etc/alertmanager/alertmanager.yml

RUN mkdir -p /alertmanager && \
    chown -R nobody:nogroup etc/alertmanager /alertmanager

USER       nobody
EXPOSE     9093
VOLUME     [ "/alertmanager" ]
WORKDIR    /alertmanager
ENTRYPOINT [ "/bin/alertmanager" ]
CMD        [ "--config.file=/etc/alertmanager/alertmanager.yml", \
             "--storage.path=/alertmanager" ]
