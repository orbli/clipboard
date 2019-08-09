
GOPATH:=$(shell go env GOPATH)
PRODUCT=clipboard
PROTO=${PRODUCT}

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. ${PRODUCT}/proto/${PROTO}.proto

.PHONY: build
build: proto
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-extldflags '-static'" -o ${PRODUCT}/${PRODUCT}-srv ${PRODUCT}/main.go ${PRODUCT}/plugin.go

.PHONY: test
test:
	go test -v ${PRODUCT}/... -cover

.PHONY: docker
docker:
	docker build ${PRODUCT}/ -t ${PRODUCT}-srv:latest
