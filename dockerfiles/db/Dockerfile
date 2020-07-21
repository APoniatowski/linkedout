# Makefile will handle compilation before proceeding with image creation.

FROM opensuse/leap:latest

LABEL MAINTAINER="Adam Poniatowski <adaml.poniatowski@gmail.com>"

# Preparation
RUN zypper refresh
RUN zypper -n install gnupg curl
RUN curl -LO https://www.mongodb.org/static/pgp/server-4.2.asc -o-
RUN rpm --import ./server-4.2.asc
RUN zypper addrepo --gpgcheck "https://repo.mongodb.org/zypper/suse/15/mongodb-org/4.2/x86_64/" mongodb
RUN zypper refresh
RUN zypper -n update

# Installation
RUN zypper -n install mongodb-org

### TODO -> add selinux for extra security later

# Create mongodb user and group
RUN groupadd mongodb
RUN useradd -g mongodb mongodb

# Create DB directory and change ownership to mongodb
RUN mkdir -p /linkedout/db
RUN chown -R mongodb:mongodb /linkedout

# Create a volume
VOLUME [ "/linkedout/db" ]

# Expose mongoDB ports
EXPOSE 27017 28017

# Run mongodb
CMD [ "exec", "su-exec", "mongodb", "mongod", "--dbpath", "/linkedout/db" ] 


# Once images are prepped, it will be passed to local registry in k8s and apply linkedout.yaml (deployment, service, etc)

####################### Basic layout, will continue with this later #############################