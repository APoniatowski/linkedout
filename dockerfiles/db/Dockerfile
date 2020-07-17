# Makefile will handle compilation before proceeding with image creation.

# mongoDB server - backend job will keep the DB updated
FROM debian:buster-slim

LABEL MAINTAINER="Adam Poniatowski <adaml.poniatowski@gmail.com>"

# Change the timezone to one you are thinking of running it in
# ENV TIMEZONE Europe/Warsaw

# Installation
# RUN echo "Installing tools..."
RUN apt-get update
RUN apt-get install gnupg wget -y
RUN wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
RUN echo "deb http://repo.mongodb.org/apt/debian buster/mongodb-org/4.2 main" | tee /etc/apt/sources.list.d/mongodb-org-4.2.list
# RUN echo "Updating..."
RUN apt-get update && apt-get upgrade -y && apt-get dist-upgrade -y
# RUN echo "Installing MongoDB..."
RUN apt-get install mongodb-org -y
# RUN echo "Removing tools..."
RUN apt-get remove gnupg wget -y

RUN mkdir -p /linkedout/db

VOLUME [ "/linkedout/db" ]

RUN chown -R mongodb:mongodb /linkedout

COPY db-init.sh .
RUN db-init.sh

EXPOSE 27017 28017

CMD [ "exec", "su-exec", "mongodb", "mongod", "--dbpath", "/linkedout/db" ] 

# Once images are prepped, it will be passed to local registry in k8s and apply linkedout.yaml (deployment, service, etc)


####################### Basic layout, will continue with this later #############################