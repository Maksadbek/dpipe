FROM busybox:ubuntu-14.04

MAINTAINER Maksadbek Akxmedov "<a.maksadbek@gmail.com>"

WORKDIR /app

COPY dpipe /app/

COPY config.toml /app/config.toml

VOLUME ["/app/data"]

ENTRYPOINT ["/app/dpipe", "-config", "/app/config.toml"]
