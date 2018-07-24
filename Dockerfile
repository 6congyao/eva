FROM alpine:3.7
LABEL MAINTAINER "6congyao@gmail.com"

RUN apk add --no-cache bash ca-certificates wget
RUN wget -q -O /etc/apk/keys/sgerrand.rsa.pub https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub
RUN wget https://github.com/sgerrand/alpine-pkg-glibc/releases/download/2.27-r0/glibc-2.27-r0.apk
RUN apk add glibc-2.27-r0.apk

ADD bin/alpine/evasvc /bin/

EXPOSE 8080

HEALTHCHECK CMD ["/bin/evasvc", "ping"]

ENTRYPOINT ["/bin/evasvc"]