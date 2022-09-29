#基于go镜像
#FROM golang:latest
FROM golang:1.19 AS build-env

#维护人的信息
MAINTAINER The Iotex Project <jeson@imoocc.com.com>

#开启80端口
EXPOSE 8888

#创建代码目录
RUN mkdir -p /w3bstream

#复制代码文件至镜像中web站点下
COPY . /w3bstream/

#Nginx配置文件软链接
#RUN ln -s /mysite/conf/test.conf /etc/nginx/conf.d/test_conf.conf
WORKDIR /w3bstream
#安装依赖

RUN cd cmd/srv-applet-mgr && go build -mod vendor
RUN mkdir -p build
RUN mv cmd/srv-applet-mgr/srv-applet-mgr build
RUN cp -r cmd/srv-applet-mgr/config build/config
RUN echo 'succeed! srv-applet-mgr =>build/srv-applet-mgr*'
RUN echo 'succeed! config =>build/config/'
RUN echo 'modify config/local.yaml to use your server config'

#migrate
RUN go run cmd/srv-applet-mgr/main.go migrate

#

#复制该脚本至镜像中，并修改其权限
ADD run.sh /run.sh
RUN chmod 775 /run.sh

#当启动容器时执行的脚本文件
CMD ["/run.sh"]