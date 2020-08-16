# Go parameters

# CGO_ENABLED=0   -> Disable interoperate with C libraries -> speed up build time! Enable it, if dependencies use C libraries!
# GOOS=linux      -> compile to linux because scratch docker file is linux
# GOARCH=amd64    -> because, hmm, everthing works fine with 64 bit :)
# -a              -> force rebuilding of packages that are already up-to-date.

BINARY_NAME=tp-link-hs110-api
BINARY_AMD64=$(BINARY_NAME)_amd64
BINARY_ARM=$(BINARY_NAME)_arm
BINARY_WINDOWS=$(BINARY_NAME)_windows.exe

all: clean test build

build: deps
	@echo "Build executable"
	@go build -a -o bin/$(BINARY_NAME)
	@echo "Done"

test:
	@echo "Run tests"
	go test -v ./...
	@echo "Done"

clean:
	@echo "Clean up go and bin directory"
	@go clean
	@rm -f bin/$(BINARY_NAME)
	@rm -f bin/$(BINARY_AMD64)
	@rm -f bin/$(BINARY_ARM)
	@rm -f bin/$(BINARY_WINDOWS)
	@rmdir bin
	@echo "Done"

deps:
	@echo "Get dependencies"
	@go get
	@echo "Done"

compile:
	@echo "Compiling for every OS and Platform"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/$(BINARY_AMD64)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 go build -a -o bin/$(BINARY_ARM)
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -o bin/$(BINARY_WINDOWS)
