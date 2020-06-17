# Makefile will handle compilation before proceeding with image creation. 

# API server - REST API to mongodb server
FROM golang:alpine
LABEL MAINTAINER="Adam Poniatowski <adaml.poniatowski@gmail.com>"

RUN mkdir -p /app/
COPY api-microservice/ /app/

WORKDIR /app
RUN go build
RUN chmod +x api-microservice

EXPOSE 80
CMD [ "/app/api-microservice" ] 

# Once images are prepped, it will be passed to local registry in k8s and apply linkedout.yaml (deployment, service, etc)


####################### Basic layout, will continue with this later #############################