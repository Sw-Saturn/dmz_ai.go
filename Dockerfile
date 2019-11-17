FROM golang:1.13.3-alpine

RUN apk add --update --no-cache build-base
ENV build_deps 'git curl bash openssl'
ENV MECAB_DIC_PATH "/usr/lib/mecab/dic/mecab-ipadic-neologd"
ENV MECAB_WORK=/tmp/mecab
ENV NEOLOGD_WORK=/tmp/neologd

RUN set -x && apk add --no-cache ${build_deps} && mkdir -p ${MECAB_DIC_PATH}

# Install MeCab
WORKDIR ${MECAB_WORK}
RUN git clone https://github.com/taku910/mecab.git ${MECAB_WORK} \
  && cd ./mecab \
  && ./configure --enable-utf8-only --with-charset=utf8 && make && make install

# Install Neologd
WORKDIR ${NEOLOGD_WORK}
RUN git clone --depth 1 https://github.com/neologd/mecab-ipadic-neologd.git ${NEOLOGD_WORK} \
  && ${NEOLOGD_WORK}/bin/install-mecab-ipadic-neologd -n -y -p ${MECAB_DIC_PATH}

WORKDIR /go/src/github.com/Sw-Saturn/dmz_ai.go
RUN apk del --purge ${build_deps} \
  && rm -rf ${MECAB_WORK} && rm -rf ${NEOLOGD_WORK}
COPY . .

ENV CGO_CFLAGS "-I/usr/include"
ENV CGO_LDFLAGS "-L/usr/lib -lmecab -lstdc++"
ENV CGO_ENABLED 1
RUN ["go","run","main.go"]