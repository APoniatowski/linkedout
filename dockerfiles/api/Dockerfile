# Makefile will handle compilation before proceeding with image creation. 

# API server - REST API to mongodb server
FROM alpine:latest
COPY ../api/bin/file .
RUN chmod +x file
RUN file 

# Once images are prepped, it will be passed to local registry in k8s and apply linkedout.yaml (deployment, service, etc)


####################### Basic layout, will continue with this later #############################