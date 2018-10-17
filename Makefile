# Go パラメータ
GOOS=linux
GOARCH=amd64
GOCMD=go


dep:
	dep ensure
dep-update:
	dep ensure -update
	dep ensure
	dep status

build:
	go build -o bin/n0core -v cmd/n0core/*.go
build-docker:
	docker build -t n0stack/n0core .
build-proto:
	docker run -it --rm  -v $(PWD):/src:ro -v `go env GOPATH`/src:/dst n0stack/build-proto --go_out=plugins=grpc:/dst

up: build-docker
	mkdir -p sandbox
	docker-compose up -d --scale mock_agent=0
up-mock: build-docker
	mkdir -p sandbox
	docker-compose up -d
logs:
	docker-compose logs -f
rm:
	docker-compose down
	docker-compose rm
	sudo rm -rf sandbox/etcd
clean:
	go clean
	rm -rf bin

analysis:
	gofmt -d -s `find ./ -name "*.go" | grep -v vendor`
	golint ./... | grep -v vendor # https://github.com/golang/lint/issues/320

test-small:
	go test -cover ./...
test-small-v:
	go test -v -cover ./...
test-small-docker:
	docker run -it --rm -v $(PWD):/go/src/github.com/n0stack/n0core n0stack/n0core make test-small

test-medium: up-mock # with root, having dependency for external
	sudo go test -tags=medium -cover ./...
test-medium-v: up-mock
	sudo go test -tags=medium -v -cover ./...
test-medium-without-root: up-mock
	go test -tags="medium without_root" -cover ./...
test-medium-without-external:
	sudo go test -tags="medium without_external" -cover ./...

run-all-in-one: up
	docker run --rm -it -v $(PWD)/bin:/go/src/github.com/n0stack/n0core/bin n0stack/n0core make build
	sudo ./bin/n0core agent \
		--name=test \
		--advertise-address=10.20.180.1 \
		--node-api-endpoint=localhost:20180 \
		--base-directory=./sandbox/workdir
