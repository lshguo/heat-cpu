FROM        debian:stretch-slim
MAINTAINER  GuoLiangShuai


COPY heatcpu  /bin/heatcpu

ENTRYPOINT [ "/bin/heatcpu" ]
