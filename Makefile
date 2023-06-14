CC=clang++ -std=c++17 -stdlib=libc++
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOFILES= ./src/go/*.go
CPPFILES= ./src/cpp/*.cpp
BINARY_NAME=GoFiler

.PHONY: all go cpp clean go-build go-clean cpp-build cpp-clean

all: go cpp

go: go-build

cpp: cpp-build

go-build:
	$(GOBUILD) -o $(BINARY_NAME) $(GOFILES)

go-clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

cpp-build:
	$(CC) -o $(BINARY_NAME)_cpp $(CPPFILES)

cpp-clean:
	rm -f $(BINARY_NAME)_cpp

clean: go-clean cpp-clean