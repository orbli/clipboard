
GOPATH:=$(shell go env GOPATH)
PRODUCT=token

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/${PRODUCT}.proto

.PHONY: build
build: proto
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-extldflags '-static'" -o ${PRODUCT}-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t ${PRODUCT}-srv:latest
