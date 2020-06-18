# Makefile will handle compilation before proceeding with image creation.

# Backend job - will get data from linkedin periodically and place them in a mongodb server
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
