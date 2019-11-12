# Frage: Was könnte man verbessern?
################################################################################
# BUILD STAGE
################################################################################

# GENERAL INFORMATION (FROM, LABEL)
FROM docker.io/golang:1.12.13
LABEL maintainer="jonas.roeckl@noris.de"
# BUILD TIME ARGUMENTS (ARG)
# RUNTIME ARGUMENTS (ENV)

# IMAGE SETUP (ADD, RUN, COPY, WORKDIR)
WORKDIR /go/src/github.com/noris/docker/
RUN go get -d -v github.com/labstack/echo/...

COPY server.go .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /binary .

################################################################################
# MAIN STAGE
################################################################################

# GENERAL INFORMATION (FROM, LABEL)
FROM docker.io/ubuntu:18.04
LABEL maintainer="jonas.roeckl@noris.de"
# BUILD TIME ARGUMENTS (ARG)
ARG UID=814
ARG UNAME="server"
ARG GID=814
ARG GNAME="server"

# RUNTIME ARGUMENTS (ENV)

# IMAGE SETUP (ADD, RUN, COPY, WORKDIR)
# Create non-root user to use to run the container payload.
# Command looks different on Alpine!
RUN set -x \
&& addgroup --system --gid ${GID} ${GNAME} \
&& adduser --system --disabled-login --ingroup ${GNAME} \
--no-create-home --home /nonexistent --shell /bin/false \
--uid ${UID} ${UNAME}

COPY --from=0 /binary /usr/local/bin/server

USER ${UNAME}:${GNAME}
ENTRYPOINT ["/usr/local/bin/server"]
