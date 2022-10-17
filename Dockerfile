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

RUN apk add --no-cache curl postgresql postgresql-client postgis mosquitto mosquitto-clients nginx

RUN curl -fsSL "https://github.com/pnpm/pnpm/releases/latest/download/pnpm-linuxstatic-x64" -o /bin/pnpm
RUN chmod +x /bin/pnpm

WORKDIR /w3bstream

COPY --from=build-go /w3bstream/cmd/srv-applet-mgr/srv-applet-mgr /w3bstream/srv-applet-mgr
COPY --from=build-go /w3bstream/cmd/srv-applet-mgr/config /w3bstream/config


COPY --from=build-nodejs /w3bstream/frontend /w3bstream/frontend
COPY --from=build-nodejs /w3bstream/frontend/.env.tmpl /w3bstream/frontend/.env

COPY entrypoint.sh /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]