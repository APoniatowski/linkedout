# Makefile will handle compilation before proceeding with image creation.

# Frontend server - Template generator and connection to API
FROM alpine:latest
LABEL MAINTAINER="Adam Poniatowski <adaml.poniatowski@gmail.com>"
LABEL MICROSERVICE="frontend"

RUN mkdir -p /app/
COPY frontend-microservice/ /app/

WORKDIR /app
RUN chmod +x frontend-microservice
RUN rm *.go go.*

EXPOSE 80
CMD [ "/app/frontend-microservice" ]
