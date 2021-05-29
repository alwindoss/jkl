
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_LOC=bin
JKL_BINARY_NAME=jkl
PROJECT_HOME=$(shell pwd)

all: get test build
build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./$(BINARY_LOC)/$(JKL_BINARY_NAME) -v ./cmd/$(JKL_BINARY_NAME)/...
test: 
	$(GOTEST) -v ./...
get: 
	$(GOGET) -v ./...
clean: 
	$(GOCLEAN)
	rm -rf $(BINARY_LOC)
run: build
	./$(BINARY_LOC)/$(JKL_BINARY_NAME)
