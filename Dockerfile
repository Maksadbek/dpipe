FROM busybox:ubuntu-14.04

MAINTAINER Maksadbek Akxmedov "<a.maksadbek@gmail.com>"

WORKDIR /app

COPY dpipe /app/

COPY config.toml /app/

VOLUME ["/app/data"]

ENTRYPOINT ["/app/dpipe", "-config", "config"]
