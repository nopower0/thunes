CURRENT=$(shell basename $(shell pwd))

LDFLAGS += -X "v/$(CURRENT)/util.CommitID=$(shell git rev-parse master)"
LDFLAGS += -X "v/$(CURRENT)/util.Date=$(shell date +%FT%T%z)"
LDFLAGS += -X "v/$(CURRENT)/util.Tag=$(shell git describe --always --tags)"
LDFLAGS += -X "v/$(CURRENT)/util.Branch=$(shell git rev-parse --abbrev-ref HEAD)"

http_server:
	@echo make http_server
	@go build -o $(CURRENT)_http_server cmd/http_server/main.go
.PHONY: http_server

linux_http_server: export GOOS = linux
linux_http_server:
	@echo make linux_http_server
	@go build -o $(CURRENT)_http_server_linux cmd/http_server/main.go
.PHONY: linux_http_server

image: linux_http_server
	@docker build -t nopower0/thunes .
	@docker push nopower0/thunes
.PHONY: image

