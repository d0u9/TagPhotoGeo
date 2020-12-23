GO := go
TOPDIR := $(shell realpath $$(pwd))
OUTPUT_DIR := output
BIN := tagphotogeo
SRCS := $(shell find . -type f -iname "*.go")

.PHONY: build
build: $(OUTPUT_DIR)/$(BIN)
	@echo "----- BUILD DONE -----"

$(OUTPUT_DIR)/$(BIN): $(SRCS)
	$(GO) build -o $(OUTPUT_DIR)/$(BIN) main.go

.PHONY: test
test:
	cd $(TOPDIR)/pkg/app; $(GO) test
