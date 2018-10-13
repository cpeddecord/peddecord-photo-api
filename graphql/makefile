 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=peddecord-photo-api
BINARY_UNIX=$(BINARY_NAME)_unix

build: 
				$(GOBUILD) -o $(BINARY_NAME) -v
test: 
				$(GOTEST) -v ./...
run:
				make build
				./$(BINARY_NAME)
docker-minikube:
				eval $(minikube docker-env)
				docker build . --tag mk-go-app
				docker tag mk-go-app localhost:5000/mk-go-app:latest
