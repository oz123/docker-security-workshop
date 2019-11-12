FROM ubuntu:18.04
LABEL maintainer="jonas.roeckl@noris.de"


COPY server /usr/local/bin/server
RUN set -x && echo "test" > /usr/local/bin/testfile

WORKDIR "/usr/local/bin"
ENTRYPOINT ["server"]
