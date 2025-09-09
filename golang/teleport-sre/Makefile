IMAGE ?= your-dockerhub-user/teleport-sre:latest

build:
	go build -o teleport-sre ./cmd/server

test:
	go test ./...

docker-build:
	docker build -t $(IMAGE) .

docker-push:
	docker push $(IMAGE)