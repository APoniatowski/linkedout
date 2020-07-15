# Makefile will handle compilation before proceeding with image creation.

# mongoDB server - backend job will keep the DB updated
FROM alpine:latest

LABEL MAINTAINER="Adam Poniatowski <adaml.poniatowski@gmail.com>"

# Change the timezone to one you are thinking of running it in
# ENV TIMEZONE Europe/Warsaw

# Installation
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/main' >> /etc/apk/repositories
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >> /etc/apk/repositories
RUN apk update
RUN apk add --no-cache mongodb su-exec
RUN mkdir -p /data/db

VOLUME [ "/data/db" ]

RUN chown -R mongodb:mongodb /data

EXPOSE 27017 28017

CMD [ "exec su-exec mongodb mongod --dbpath /data/db" ] 

# Once images are prepped, it will be passed to local registry in k8s and apply linkedout.yaml (deployment, service, etc)


####################### Basic layout, will continue with this later #############################