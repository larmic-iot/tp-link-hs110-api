# tp-link-hs110-api

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Build Status](https://travis-ci.com/larmic/tp-link-hs110-api.svg?branch=master)](https://travis-ci.com/larmic/tp-link-hs110-api)

## Requirements

* Docker 
* Go 1.15.x (if you want to build it without using docker builder)

## Build it

```sh 
$ make docker-build                          # build local docker image
$ make docker-push                           # push local docker image to hub.docker.com
$ make docker-all                            # build and push docker image to hub.docker.com
$ make IMAGE_TAG="0.0.1" docker-all          # build and push docker image with specific version
```

## Run it native

```sh 
$ make run                                   # start native app 
$ curl http://localhost:8080/api/10.0.0.210  # call rest service
$ ctrl+c                                     # stop native app
```

## Run it using docker

```sh 
$ make docker-run                            # start docker image 
$ curl http://localhost:8080/api/10.0.0.210  # call rest service
$ make docker-stop                           # stop and remove docker app
```