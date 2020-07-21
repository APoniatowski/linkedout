Gather_facts:
	echo "TODO: get data from user, to change services (frontend/backend) to different ports,"
	echo "      password for the backend service, if the user has a local/remote container registry,etc."

Go_microservice_binaries:
	echo "Compiling linkeout binaries..."
	cd backend-microservice; GOOS=linux GOARCH=amd64 go build
	cd api-microservice; GOOS=linux GOARCH=amd64 go build
	cd frontend-microservice; GOOS=linux GOARCH=amd64 go build
	echo "Binaries are compiled."

Docker_images:
	echo "Creating docker images..."
	docker build . -f dockerfiles/db/Dockerfile -t linkedout/mongodb:1.0.0 --network=host
	docker build . -f dockerfiles/backend/Dockerfile -t linkedout/backend-microservice:1.0.0 --network=host
	docker build . -f dockerfiles/api/Dockerfile -t linkedout/api-microservice:1.0.0 --network=host
	docker build . -f dockerfiles/frontend/Dockerfile -t linkedout/frontend-microservice:1.0.0 --network=host
	echo "Docker images have been created."

Docker_push:
	echo "TODO: will push to registry (docker/openshift/local/etc)"

K8s_deployment:
	echo "TODO: will create/apply the relevant deployments/services/secrets/configmaps/etc"

Post_deployment:
	echo "TODO: create database and collections, cleanup, etc."