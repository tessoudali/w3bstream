#build backend
FROM golang:1.19 AS build-go


COPY . /w3bstream/

#Build vendor
WORKDIR /w3bstream
RUN cd cmd/srv-applet-mgr && CGO_ENABLED=0 go build -mod vendor
RUN mkdir -p build

#Build frontend
FROM node:14.20-alpine AS build-nodejs

WORKDIR /w3bstream

COPY . /w3bstream/

RUN apk add --no-cache curl
RUN curl -fsSL "https://github.com/pnpm/pnpm/releases/latest/download/pnpm-linuxstatic-x64" -o /bin/pnpm
RUN chmod +x /bin/pnpm
RUN cd /w3bstream/frontend && pnpm install
RUN cp /w3bstream/frontend/.env.tmpl /w3bstream/frontend/.env
RUN cd /w3bstream/frontend && pnpm build


#run
FROM node:14.20-alpine

EXPOSE 8888
EXPOSE 1883
EXPOSE 3000 
EXPOSE 5432

#Add VOLUMEs to allow backup of config, logs and databases
RUN apk add --no-cache curl postgresql postgresql-client postgis mosquitto mosquitto-clients nginx

##Initialize PostgreSQL database
RUN mkdir -p /var/lib/postgresql/data && chmod 700 /var/lib/postgresql/data && chown -R postgres:postgres /var/lib/postgresql/data
RUN mkdir -p /run/postgresql/ && chown -R postgres:postgres /run/postgresql/ && chmod 777 /var/lib/postgresql
RUN su - postgres -c "initdb /var/lib/postgresql/data"
RUN echo "host all  all    0.0.0.0/0  md5" >> /var/lib/postgresql/data/pg_hba.conf
RUN su - postgres -c "pg_ctl start -D /var/lib/postgresql/data -l /var/lib/postgresql/log.log && createuser test_user && psql --command \"ALTER USER test_user WITH ENCRYPTED PASSWORD 'test_passwd';\" && psql --command \"CREATE DATABASE test;\" && psql --command \"GRANT ALL PRIVILEGES ON DATABASE test to test_user;\""
#
##Initialize MQTT
RUN mkdir /var/run/mosquitto -p
RUN rm -f /etc/mosquitto/mosquitto.conf 
COPY build_image/conf/mosquitto.conf /etc/mosquitto/mosquitto.conf 
#RUN service mosquitto start
#
#Install pnpm
RUN curl -fsSL "https://github.com/pnpm/pnpm/releases/latest/download/pnpm-linuxstatic-x64" -o /bin/pnpm
RUN chmod +x /bin/pnpm

#WORKDIR /w3bstream
COPY --from=build-go /w3bstream/cmd/srv-applet-mgr/srv-applet-mgr /w3bstream/srv-applet-mgr
COPY --from=build-go /w3bstream/cmd/srv-applet-mgr/config /w3bstream/config


COPY --from=build-nodejs /w3bstream/frontend /w3bstream/frontend
COPY --from=build-nodejs /w3bstream/frontend/.env.tmpl /w3bstream/frontend/.env

#Upload init script
ADD build_image/cmd/docker_init.sh /init.sh
RUN chmod 775 /init.sh
