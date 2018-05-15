FROM alpine:3.7
MAINTAINER Lucas <6congyao@gmail.com>

RUN apk add --no-cache bash ca-certificates wget
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.27-r0/glibc-2.27-r0.apk
RUN apk add glibc-2.27-r0.apk

ADD eva /bin/

ENV EVA_DB_DRIVER postgres
ENV EVA_DB_SOURCE postgres://postgres:root@139.198.177.115:5432/iam?sslmode=disable

EXPOSE 8080

HEALTHCHECK CMD ["/bin/eva", "ping"]

ENTRYPOINT ["/bin/eva"]