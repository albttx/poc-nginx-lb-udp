FROM yyyar/gobetween

MAINTAINER Albert Le Batteux <albttx@scw-mbp-albttx.local>

FROM ubuntu:latest

RUN apt-get update \
    && apt-get install -y iputils-ping

COPY --from=0 /usr/bin/gobetween /usr/bin/gobetween

CMD ["/usr/bin/gobetween", "-c", "/etc/gobetween/conf/gobetween.toml"]
