# list of packages without vendor
PKGS=$(shell go list ./... | grep -v "vendor")

# default dpipe build
default: prepare build

# get all dependency packages
prepare:
	go get github.com/tools/godep
	godep restore

# build and install into GOBIN directory
build:
	go install ./...

# build with godep
build-with-godep:
	go get github.com/tools/godep
	godep go build ./cmd/dpipe/dpipe.go

build-for-docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dpipe ./cmd/dpipe/dpipe.go

build-docker-image:
	docker build -t dpipe .

test: vet
	go test $(PKGS)

vet:
	go vet $(PKGS)
