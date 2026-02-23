FROM docker.io/golang:latest as builder

WORKDIR /opt/my-server

COPY . .

RUN go build . && pwd

FROM registry.access.redhat.com/ubi10

WORKDIR /opt/lightweight-server

RUN groupadd project -g 1200 && \
    useradd org -m -d /opt/org -u 1200 -g 1200

COPY --from=builder --chown=org:project  /opt/my-server/go-http-server  /opt/org

WORKDIR /opt/org

USER org

CMD [ "./go-http-server" ]
