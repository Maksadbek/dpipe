# default dpipe build
default: prepare build

# get all dependency packages
prepare:
	go get github.com/tools/godep
	godep restore

# build and install into GOBIN directory
build:
	go install ./...

build-for-docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -o dpipe -ldflags ./cmd/dpipe/dpipe.go

build-docker-image:
	docker build . -t dpipe

test: vet
	go test ./...

vet:
	go vet ./...
