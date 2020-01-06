BINDIR			:=	$(CURDIR)/bin
GOCMD				=	go
GOBUILD			=	$(GOCMD) build
GOCLEAN			=	$(GOCMD) clean
GOTEST			=	$(GOCMD) test
BINNAME			?=	jarvis

#------
#	all
.PHONY: all

all: test build

#------
#	build
.PHONY: build

build:
				$(GOBUILD) -o $(BINDIR)/$(BINNAME) -v ./cmd/jarvis

#------
#	test
.PHONY: test

test:
				$(GOTEST) -v ./...

#------
#	test
.PHONY: clean

clean:
				$(GOCLEAN)
