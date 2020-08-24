CONTAINER_NAME=tp-link-hs110-api
IMAGE_NAME=larmic/tp-link-hs110-api
IMAGE_TAG=latest

run:
	go run main.go

docker-all: docker-build docker-push

docker-build:
	@echo "Remove docker image if already exists"
	docker rmi -f ${IMAGE_NAME}:${IMAGE_TAG}
	@echo "Build go docker image"
	docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
	@echo "Prune intermediate images"
	docker image prune --filter label=stage=intermediate -f

docker-push:
	docker push ${IMAGE_NAME}:${IMAGE_TAG}

docker-run:
	docker run -d -p 8080:8080 --rm --name --rm --name ${CONTAINER_NAME} ${IMAGE_NAME}

docker-stop:
	docker stop ${CONTAINER_NAME}
