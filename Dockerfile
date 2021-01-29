# Use multi stage build to# minimize generated docker images size
# see: https://docs.docker.com/develop/develop-images/multistage-build/

# Step 1: create multi stage builder (about 800 MB)
FROM golang:1.15 AS builder
LABEL stage=intermediate
RUN go version

WORKDIR /go/src/larmic/

COPY main.go go.mod go.sum /go/src/larmic/
COPY api /go/src/larmic/api

RUN go mod download

RUN go test -v ./...

# CGO_ENABLED=0   -> Disable interoperate with C libraries -> speed up build time! Enable it, if dependencies use C libraries!
# GOOS=linux      -> compile to linux because scratch docker file is linux
# GOARCH=amd64    -> because, hmm, everthing works fine with 64 bit :)
# -a              -> force rebuilding of packages that are already up-to-date.
# -o app          -> force to build an executable app file (instead of default https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)
ARG CGO_ENABLED=0
ARG GOARCH=amd64
ARG GOARM=7
ARG TARGETPLATFORM=linux/amd64
ARG BUILDPLATFORM
RUN echo "Hello, my CPU architecture is $(uname -m)"
RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"

RUN if [ "$TARGETPLATFORM" = "linux/arm/v7" ] ; then echo "arm v7"; else ; fi
RUN if [ "$TARGETPLATFORM" = "linux/arm64" ] ; then echo "arm 64"; else ; fi
RUN if [ "$TARGETPLATFORM" = "linux/amd64" ] ; then echo "amd 64"; else ; fi

RUN env CGO_ENABLED=${CGO_ENABLED} GOARCH=${GOARCH} GOARM=${GOARM} go build -a -o main .

# Step 2: create minimal executable image (less than 10 MB)
FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/larmic/main .
COPY open-api-3.yaml .

EXPOSE 8080
ENTRYPOINT ["./main"]
