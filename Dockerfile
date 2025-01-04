FROM golang:1.23

RUN set -xe && \
    apt-get update && \
    apt-get install -y --no-install-recommends \
    build-essential \
    lsb-release wget software-properties-common gnup

RUN set -xe && \
    wget https://apt.llvm.org/llvm.sh && \
    chmod +x llvm.sh && \
    ./llvm.sh 18

WORKDIR /cellscript

CMD bash