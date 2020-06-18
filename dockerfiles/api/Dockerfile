# Makefile will handle compilation before proceeding with image creation.

# API server - REST API to mongodb server
FROM alpine:latest
LABEL MAINTAINER="Adam Poniatowski <adaml.poniatowski@gmail.com>"
LABEL MICROSERVICE="api"

RUN mkdir -p /app/
COPY api-microservice/ /app/

WORKDIR /app
RUN chmod +x api-microservice
RUN rm *.go go.*

EXPOSE 80
CMD [ "/app/api-microservice" ]