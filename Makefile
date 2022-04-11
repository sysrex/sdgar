.PHONY: clean critic security lint test build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build


clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

docker.run: docker.network docker.mariadb swag docker.fiber docker.redis

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name sdgar-fiber \
		--network dev-network \
		-p 8000:8000 \
		fiber

docker.mariadb:
	docker run --rm -d \
		--name sdgar-mariadb \
		--network dev-network \
		-e MYSQL_ROOT_PASSWORD=mysql \
		-e MYSQL_USER=sdgar \
		-e MYSQL_PASSWORD=sdgar \
		-e MYSQL_DATABASE=sdgar \
		-p 3306:3306 \
		-d docker.io/library/mariadb:10

docker.redis:
	docker run --rm -d \
		--name sdgar-redis \
		--network dev-network \
		-p 6379:6379 \
		redis

docker.stop: docker.stop.fiber docker.stop.mariadb docker.stop.redis

docker.stop.fiber:
	docker stop sdgar-fiber

docker.stop.postgres:
	docker stop sdgar-mariadb

docker.stop.redis:
	docker stop sdgar-redis

swag:
	swag init