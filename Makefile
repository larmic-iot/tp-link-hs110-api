# Go parameters

# go get
# go test -v ./...
# CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o tp-link-hs110-api
# go clean

# CGO_ENABLED=0   -> Disable interoperate with C libraries -> speed up build time! Enable it, if dependencies use C libraries!
# GOOS=linux      -> compile to linux because scratch docker file is linux
# GOARCH=amd64    -> because, hmm, everthing works fine with 64 bit :)
# -a              -> force rebuilding of packages that are already up-to-date.

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
