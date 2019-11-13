FROM alpine:3.9

RUN apk add --no-cache mariadb

COPY run.sh /run.sh

COPY my.cnf /etc/

VOLUME ["/var/lib/mysql"]
ENTRYPOINT ["/run.sh"]
EXPOSE 3306
