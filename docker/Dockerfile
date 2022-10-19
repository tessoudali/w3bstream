#build backend
FROM golang:1.19 AS build-go

WORKDIR /w3bstream

COPY . /w3bstream/
RUN make build_server


#build frontend
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
EXPOSE 80

RUN apk add --no-cache curl postgresql postgresql-client postgis mosquitto mosquitto-clients nginx openrc

RUN curl -fsSL "https://github.com/pnpm/pnpm/releases/latest/download/pnpm-linuxstatic-x64" -o /bin/pnpm
RUN chmod +x /bin/pnpm

WORKDIR /w3bstream

COPY --from=build-go /w3bstream/cmd/srv-applet-mgr/build/srv-applet-mgr /w3bstream/build/srv-applet-mgr
COPY --from=build-go /w3bstream/cmd/srv-applet-mgr/build/config /w3bstream/build/config


COPY --from=build-nodejs /w3bstream/frontend /w3bstream/frontend
COPY --from=build-nodejs /w3bstream/frontend/.env.tmpl /w3bstream/frontend/.env

#added to runlevel sysinit
RUN rc-update add nginx
RUN rc-update add postgresql
RUN rc-update add mosquitto

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod a+x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]