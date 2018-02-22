FROM        debian:stretch-slim
MAINTAINER  GuoLiangShuai


COPY heat-cpu  /bin/heatcpu

ENTRYPOINT [ "/bin/heatcpu" ]
