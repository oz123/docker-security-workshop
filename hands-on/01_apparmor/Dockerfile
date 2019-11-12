# Frage: Was ist hier alles schlecht?

FROM ubuntu:latest
LABEL maintainer="jonas.roeckl@noris.de"

RUN apt-get update && apt-get install -y iputils-ping curl

COPY server /usr/local/bin/server

ENTRYPOINT ["/usr/local/bin/server"]
