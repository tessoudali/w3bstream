#Build go vendor
FROM golang:1.19 AS build-go
COPY . /w3bstream/

WORKDIR /w3bstream
RUN cd cmd/srv-applet-mgr && GOWORK=off GOOS=linux CGO_ENABLED=1 go build
RUN mkdir -p build

#Build noodjs
FROM node:14.20 AS build-nodejs
#FROM node:14.20-alpine AS build-nodejs
#FROM node:16-alpine AS build-nodejs

WORKDIR /w3bstream-nodejs

#RUN apk add --no-cache curl
#RUN curl -fsSL "https://github.com/pnpm/pnpm/releases/latest/download/pnpm-linuxstatic-x64" -o /bin/pnpm
#RUN chmod +x /bin/pnpm
RUN npm i -g pnpm
COPY ./studio .
RUN npm i pnpm -g
RUN pnpm install --no-frozen-lockfile
RUN pnpm i --frozen-lockfile;
RUN pnpm build:standalone
RUN sed -i 's,"http://localhost:8888",process.env.NEXT_PUBLIC_API_URL,g' .next/standalone/server.js


#run
FROM golang:1.19

EXPOSE 5432
EXPOSE 8888
EXPOSE 1883
EXPOSE 3000 
EXPOSE 22
#Create user postgres
RUN mkdir -p /home/postgres
RUN groupadd postgres --gid=999 \
  && useradd -d /home/postgres --gid postgres --uid=999 postgres
RUN chown -R postgres:postgres /home/postgres

#Install postgresql and postgis
RUN apt -y update
RUN apt install postgresql postgresql-contrib -y
RUN apt-get install net-tools -y
RUN apt install postgis postgresql-13-postgis-3 -y

#Create app directory and upload w3bstream app
RUN mkdir -p /w3bstream
WORKDIR /w3bstream


#Initialize PostgreSQL database
RUN /etc/init.d/postgresql start && \
  su postgres sh -c "psql -c \"CREATE USER test_user WITH ENCRYPTED PASSWORD 'test_passwd'\"" && \
  su postgres sh -c "psql -c \"CREATE DATABASE test\"" && \
  su postgres sh -c "psql -c \"GRANT ALL PRIVILEGES ON DATABASE test to test_user;;\""

RUN echo "listen_addresses='*'" >> /etc/postgresql/13/main/postgresql.conf
RUN echo "host all all 0.0.0.0/0 md5" >> /etc/postgresql/13/main/pg_hba.conf
#Install mqtt
RUN apt install mosquitto mosquitto-clients -y
RUN mkdir /var/run/mosquitto -p
RUN sed -i 's,/run/mosquitto/mosquitto.pid,/tmp/mosquitto.pid,g' /etc/mosquitto/mosquitto.conf

#Install Nginx
RUN apt-get install nginx -y

#Install Nodejs and pnpm
RUN mkdir -p /w3bstream/build_image/packages
COPY build_image/packages/setup_14.x /w3bstream/build_image/packages/setup_14.x
RUN /bin/bash /w3bstream/build_image/packages/setup_14.x
RUN apt install nodejs -y
RUN mkdir -p /w3bstream/build_image/packages
COPY build_image/packages/install.sh /w3bstream/build_image/packages/install.sh
RUN /bin/bash /w3bstream/build_image/packages/install.sh
RUN ln -s /root/.local/share/pnpm/pnpm /usr/bin/pnpm

#WORKDIR /w3bstream
RUN mkdir -p /w3bstream/cmd/srv-applet-mgr/config
COPY --from=build-go /w3bstream/cmd/srv-applet-mgr/srv-applet-mgr /w3bstream/cmd/srv-applet-mgr/srv-applet-mgr
COPY --from=build-go /w3bstream/cmd/srv-applet-mgr/config/default.yml /w3bstream/cmd/srv-applet-mgr/config/default.yml

COPY --from=build-nodejs /w3bstream-nodejs/public ./studio-build/public
COPY --from=build-nodejs /w3bstream-nodejs/.next/standalone ./studio-build
COPY --from=build-nodejs /w3bstream-nodejs/.next/static ./studio-build/.next/static

#Upload init script and configure
RUN mkdir /w3bstream/build_image -p 
ADD build_image/cmd/docker_init.sh /init.sh
RUN chmod 775 /init.sh
