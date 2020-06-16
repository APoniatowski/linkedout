# Makefile will handle compilation before proceeding with image creation. 

# Backend job - will get data from linkedin periodically and place them in a mongodb server
FROM alpine:latest
LABEL authors="Adam Poniatowski"
COPY ../backend/bin/file .
RUN chmod +x file
RUN file

# Once images are prepped, it will be passed to local registry in k8s and apply linkedout.yaml (deployment, service, etc)


####################### Basic layout, will continue with this later #############################