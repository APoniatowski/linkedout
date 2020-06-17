# Makefile will handle compilation before proceeding with image creation. 

# Frontend server - Template generator and connection to API
FROM golang:alpine
LABEL MAINTAINER="Adam Poniatowski <adaml.poniatowski@gmail.com>"

RUN mkdir -p /app/
COPY frontend-microservice/ /app/

WORKDIR /app
RUN go build
RUN chmod +x frontend-microservice

EXPOSE 80
CMD [ "/app/frontend-microservice" ]

# Once images are prepped, it will be passed to local registry in k8s and apply linkedout.yaml (deployment, service, etc)


####################### Basic layout, will continue with this later #############################