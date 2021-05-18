FROM golang:1.15.10-alpine3.12

ENV MECAB_VERSION 0.996
ENV IPADIC_VERSION 2.7.0-20070801
ENV MECAB_URL https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7cENtOXlicTFaRUE
ENV IPADIC_URL https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7MWVlSDBCSXZMTXM

RUN apk update && \
    apk upgrade && \
    apk add --no-cache build-base curl git bash openssl openssh-client sudo && \
    curl -SL -o mecab-${MECAB_VERSION}.tar.gz ${MECAB_URL} && \
    tar zxf mecab-${MECAB_VERSION}.tar.gz && \
    cd mecab-${MECAB_VERSION} && \
    ./configure --enable-utf8-only --with-charset=utf8 && \
    make && \
    make install && \
    cd && \
    curl -SL -o mecab-ipadic-${IPADIC_VERSION}.tar.gz ${IPADIC_URL} && \
    tar zxf mecab-ipadic-${IPADIC_VERSION}.tar.gz && \
    cd mecab-ipadic-${IPADIC_VERSION} && \
    ./configure --with-charset=utf8 && \
    make && \
    make install && \
    cd && \
    git clone --depth 1 https://github.com/neologd/mecab-ipadic-neologd.git /mecab-ipadic-neologd && \
    /mecab-ipadic-neologd/bin/install-mecab-ipadic-neologd -n -a -y -p /usr/local/lib/mecab/dic/mecab-ipadic-neologd && \
    echo "dicdir = /usr/local/lib/mecab/dic/mecab-ipadic-neologd" > /usr/local/etc/mecabrc && \
    rm -rf \
      mecab-${MECAB_VERSION}* \
      mecab-${IPADIC_VERSION}* \
      /mecab-ipadic-neologd

WORKDIR /go/src/mecab_converter/application

COPY ./application .
RUN chmod 744 ./start.sh

EXPOSE 8080

CMD ["sh", "./start.sh"]

