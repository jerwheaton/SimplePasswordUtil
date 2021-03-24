GO := GO111MODULE=on go
MAIN_SRC_FILE := cmd/password/password.go
BUILD_TARGET := build
BUILD_DIR := build
NAME := password

all: build

build:
	$(GO) $(BUILD_TARGET) -o $(BUILD_DIR)/$(NAME) $(MAIN_SRC_FILE)

clean:
	rm -rf build out/*
