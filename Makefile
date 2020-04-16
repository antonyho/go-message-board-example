GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
TARGET=messageboard-api-example

all: build

.PHONY: api dep dep-update test benchmark run image build

api:
	docker run --rm \
	-v $(PWD):/local openapitools/openapi-generator-cli generate \
	-i /local/resources/api/spec/v1/swagger.json \
	-g go-server \
	-o /local/pkg/openapi

api-doc:
	docker run --rm \
	-v $(PWD):/local openapitools/openapi-generator-cli generate \
	-i /local/resources/api/spec/v1/swagger.json \
	-g html \
	-o /local/resources/api/doc

dep:
	go mod init

dep-update:
	go mod tidy

test:
	$(GOTEST) -v -race -cover ./...

benchmark:
	$(GOTEST) -race -bench=. ./...

build:
	$(GOBUILD) -installsuffix cgo -o $(TARGET) ./cmd/restserver

clean:
	rm $(TARGET)
	$(GOCLEAN)

image:
	docker build --rm -t antonyho-messageboard-api-example .

run:
	$(GORUN) cmd/restserver/main.go
