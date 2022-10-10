#FROM golang:latest
FROM imoocc/w3b_base:v1 AS build-env

MAINTAINER The Iotex Project <jeson@iotex.io>

EXPOSE 8888
EXPOSE 3000 
EXPOSE 5432
EXPOSE 22

RUN mkdir -p /w3bstream

COPY . /w3bstream/

#RUN ln -s /mysite/conf/test.conf /etc/nginx/conf.d/test_conf.conf
WORKDIR /w3bstream

RUN cd cmd/srv-applet-mgr && go build -mod vendor
RUN mkdir -p build
RUN mv cmd/srv-applet-mgr/srv-applet-mgr build
RUN cp -r cmd/srv-applet-mgr/config build/config
RUN echo 'succeed! srv-applet-mgr =>build/srv-applet-mgr*'
RUN echo 'succeed! config =>build/config/'
RUN echo 'modify config/local.yaml to use your server config'

#w3bstream front
#RUN go run cmd/srv-applet-mgr/main.go migrate
RUN cd /w3bstream/frontend && pnpm install
RUN cp /w3bstream/frontend/.env.tmpl /w3bstream/frontend/.env
RUN cd /w3bstream/frontend && pnpm build

#
#init script
ADD cmd/docker_init.sh /init.sh
RUN chmod 775 /init.sh
