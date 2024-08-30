include ../lib.mk
IMAGE ?= $(REPO_NS)/protoc-gen-lua-v2

bin:
	go mod tidy
	go mod download
	CGO_ENABLED=0 go build -ldflags '-s -w' -o proto-reverse-lua-v2 main.go

image: bin
	docker build . -t $(IMAGE)

push:
	docker push $(IMAGE)

docker-test:
	docker run --rm -w /usr/local/protoc-gen-lua/example $(IMAGE) sh -c 'protoc --lua_out=./ person.proto && cat person_pb.lua'

test:
	go run main.go testdata/list_pb.lua
	go run main.go testdata/person_pb.lua
