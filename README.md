# tp-link-hs110-api

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![Go](https://github.com/larmic/tp-link-hs110-api/workflows/Go/badge.svg)
[![Docker hub image](https://img.shields.io/docker/image-size/larmic/tp-link-hs110-api?label=dockerhub)](https://hub.docker.com/repository/docker/larmic/tp-link-hs110-api)

A REST api client (adapter) for the proprietary TP-Link Smart Home protocol to control TP-Link HS110 WiFi Smart Plugs. 
The SmartHome protocol runs on TCP port 9999 and uses a trivial XOR autokey encryption that provides no security.
This application uses a REST api as a wrapper on the TCP socket connection. There is no authentication mechanism and 
REST endpoints documented in [open api 3.1](open-api-3.yaml).

This project inspired by [tplink-smartplug written in python](https://github.com/softScheck/tplink-smartplug) and 
improves my Go knowledge.

## Usage

The easiest way is to use the docker image. Otherwise, the artifact will have to be built by yourself.

```sh 
$ docker pull larmic/tp-link-hs110-api
$ docker run -d -p 8080:8080 --rm --name larmic-tp-link-hs110-api larmic/tp-link-hs110-api
```

## Example requests

```sh 
$ curl http://localhost:8080                    # Open Api 3.1 specification
$ curl http://localhost:8080/10.0.0.1           # General energy plug information
$ curl http://localhost:8080/10.0.0.1/energy    # Energy consumption
```

## Build application by yourself

### Requirements

* Docker 
* Go 1.15.x (if you want to build it without using docker builder)

### Build it

```sh 
$ make docker-build                             # build local docker image
$ make docker-push                              # push local docker image to hub.docker.com
$ make docker-all                               # build and push docker image to hub.docker.com
$ make IMAGE_TAG="0.0.1" docker-all             # build and push docker image with specific version
```

### Run it native

```sh 
$ make run                                      # start native app 
$ curl http://localhost:8080/api/10.0.0.210     # call rest service
$ ctrl+c                                        # stop native app
```

### Run it using docker

```sh 
$ make docker-run                               # start docker image 
$ curl http://localhost:8080/api/10.0.0.210     # call rest service
$ make docker-stop                              # stop and remove docker app
```