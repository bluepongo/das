# Commands
GOCMD= GO111MODULE=on go
GOBUILD=	$(GOCMD) build
GOCLEAN=	$(GOCMD) clean
GOTEST=		$(GOCMD) test
GOGET=		$(GOCMD) get
GOMODULE=	$(GOCMD) mod download

# Paths
BINARY_NAME=das

all: deps test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v
build-linux: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

test:
	$(GOTEST) -v ./...

deps:
	$(GOMODULE)

.PHONY:clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)


	