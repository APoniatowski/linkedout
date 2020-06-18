# Makefile will handle compilation before proceeding with image creation.

# Backend job - an admin panel that will use the api to update the mongodb server
############### as well as making api requests to github/gitlab for projects
FROM alpine:latest
LABEL MAINTAINER="Adam Poniatowski <adaml.poniatowski@gmail.com>"
LABEL MICROSERVICE="backend"

RUN mkdir -p /app/
COPY backend-microservice/ /app/

WORKDIR /app
RUN chmod +x backend-microservice
RUN rm *.go go.*

EXPOSE 80
CMD [ "/app/backend-microservice" ]
