FROM golang:alpine

ENV app_name=go-crawler
ENV organization=ujway
ENV http_port=8080
ENV https_port=10443
ENV EDITOR=vim
ENV TZ Asia/Tokyo

# Prepare modules
RUN apk update && \
    apk upgrade && \
    apk add --update --no-cache --virtual=.build-dependencies \
      alpine-sdk \
      curl-dev \
      libxml2-dev \
      libxslt-dev \
      nodejs \
      ca-certificates \
      libstdc++ \
      yaml-dev \
      zlib-dev && \
   apk add --update --no-cache \
      ncurses \
      libxslt \
      tzdata \
      bash \
      yaml \
      vim \
      less \
      yarn \
      mysql-client \
      git \
      build-base

RUN go get github.com/PuerkitoBio/goquery
RUN mkdir -p $GOPATH/src/github.com/$organization/$app_name
COPY . $GOPATH/src/github.com/$organization/$app_name
WORKDIR $GOPATH/src/github.com/$organization/$app_name
RUN echo "export PS1='\[\e[0;32m\]\u@\h:\[\e[0;36m\]\w\[\e[0m\]$ '" >> /root/.bashrc

EXPOSE $http_port
EXPOSE $https_port
CMD tail -f /dev/null
