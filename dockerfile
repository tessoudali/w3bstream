# Golang Image
#FROM golang:latest
FROM golang:1.19 AS build-env

# Maintainer Information
MAINTAINER The Iotex Project <jeson@imoocc.com.com>

# Port 8888
EXPOSE 8888

# Create Project Directory
RUN mkdir -p /w3bstream

# Copy Files into Project Directory
COPY . /w3bstream/

# Link Nginx Config File
#RUN ln -s /mysite/conf/test.conf /etc/nginx/conf.d/test_conf.conf
WORKDIR /w3bstream

# Install Dependencies

RUN cd cmd/srv-applet-mgr && go build -mod vendor
RUN mkdir -p build
RUN mv cmd/srv-applet-mgr/srv-applet-mgr build
RUN cp -r cmd/srv-applet-mgr/config build/config
RUN echo 'succeed! srv-applet-mgr =>build/srv-applet-mgr*'
RUN echo 'succeed! config =>build/config/'
RUN echo 'modify config/local.yaml to use your server config'

# Run DB Migrate
RUN go run cmd/srv-applet-mgr/main.go migrate

# Copy Script into Image and Modify Permission
ADD run.sh /run.sh
RUN chmod 775 /run.sh

# Docker Container Starting Command
CMD ["/run.sh"]
