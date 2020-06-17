# Makefile will handle compilation before proceeding with image creation. 

# Backend job - will get data from linkedin periodically and place them in a mongodb server
FROM golang:alpine
LABEL MAINTAINER="Adam Poniatowski <adaml.poniatowski@gmail.com>"

RUN mkdir -p /app/
COPY backend-microservice/ /app/

WORKDIR /app
RUN go build
RUN chmod +x backend-microservice

EXPOSE 80
CMD [ "/app/backend-microservice" ]

# Once images are prepped, it will be passed to local registry in k8s and apply linkedout.yaml (deployment, service, etc)


####################### Basic layout, will continue with this later #############################