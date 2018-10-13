 # Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=peddecord-photo-api
BINARY_UNIX=$(BINARY_NAME)_unix

api-get:
				(cd graphql;$(GOGET) -d -v)
api-build: 
				make api-get
				(cd graphql;$(GOBUILD) -o ../$(BINARY_NAME) -v)
api-build-ci:
				make api-get
				(cd graphql;CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix -o ../$(BINARY_NAME))
api-test:
				(cd graphql;$(GOTEST))
api-run:
				make api-build
				./$(BINARY_NAME)
				rm ./$(BINARY_NAME)
docker-minikube:
				eval $(minikube docker-env)
				docker build . --tag mk-go-app
				docker tag mk-go-app localhost:5000/mk-go-app:latest
