# Makefile will handle compilation before proceeding with image creation. 

# Frontend server - Template generator and connection to API
FROM alpine:latest
COPY ../frontend/bin/file .
RUN chmod +x file
RUN file
EXPOSE 8080

# Once images are prepped, it will be passed to local registry in k8s and apply linkedout.yaml (deployment, service, etc)


####################### Basic layout, will continue with this later #############################