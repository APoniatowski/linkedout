Go_microservice_binaries:
	cd backend-microservice; go build
	cd api-microservice; go build
	cd frontend-microservice; go build

Docker_images:
	docker build . -f dockerfiles/db/Dockerfile -t linkedout/mongodb:1.0.0 --network=host
	docker build . -f dockerfiles/backend/Dockerfile -t linkedout/backend-microservice:1.0.0 --network=host
	docker build . -f dockerfiles/api/Dockerfile -t linkedout/api-microservice:1.0.0 --network=host
	docker build . -f dockerfiles/frontend/Dockerfile -t linkedout/frontend-microservice:1.0.0 --network=host

Docker_push:
	echo "TODO: will push to registry (docker/openshift/local/etc)"

K8s_deployment:
	echo "TODO: will create/apply the relevant deployments/services/secrets/configmaps/etc"