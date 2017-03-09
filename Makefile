# default dpipe build
default: prepare build

# get all dependency packages
prepare:
	go get github.com/tools/godep
	godep restore

# build and install into GOBIN directory
build:
	go install ./...

test: vet
	go test ./...

vet:
	go vet ./...
