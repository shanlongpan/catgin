FROM centos:centos7

LABEL MAINTAINER=shanlongpan

COPY catgin /data/

ENTRYPOINT ["/data/catgin"]