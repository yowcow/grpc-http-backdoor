FROM yowcow/debian-buster-perl5:5.28.0

RUN set -eux; \
    apt-get update && \
    apt-get -yq install \
        libprotobuf-dev \
        libprotoc-dev \
        libssl-dev \
        protobuf-compiler

RUN set -eux; \
    cpanm --notest Carton

RUN set -eux; \
    rm -rf /root/.cpanm /var/lib/apt/lists/*
