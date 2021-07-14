PROTO_PATH=./src
SRC = $(shell find src -iname "*.proto" | grep -v src/google) src/google/api/annotations.proto src/google/api/http.proto

GEN_SRC_ROOT = ./gen_src
GEN_SRC_CSHARP = $(GEN_SRC_ROOT)/csharp
GEN_SRC_CPP = $(GEN_SRC_ROOT)/cpp
GEN_SRC_JAVA = $(GEN_SRC_ROOT)/java
GEN_SRC_JS = $(GEN_SRC_ROOT)/js
GEN_SRC_TS = $(GEN_SRC_ROOT)/ts
GEN_SRC_PHP = $(GEN_SRC_ROOT)/php
GEN_SRC_PYTHON = $(GEN_SRC_ROOT)/python
GEN_SRC_DART = $(GEN_SRC_ROOT)/dart
GEN_SRC_GO = $(GEN_SRC_ROOT)/go
GEN_SRC_RUST = $(GEN_SRC_ROOT)/rust
GEN_SRC_RUBY = $(GEN_SRC_ROOT)/ruby

.PHONY: disable_default
disable_default:
	$(info Please specify target.)

.PHONY: all
all: clean csharp cpp java js php python go

.PHONY: csharp
csharp:
	rm -rf ${GEN_SRC_CSHARP}
	mkdir -p $(GEN_SRC_CSHARP)
	protoc --proto_path=$(PROTO_PATH) --csharp_out=$(GEN_SRC_CSHARP) --csharp_opt=base_namespace $(SRC)

.PHONY: cpp
cpp:
	rm -rf ${GEN_SRC_CPP}
	mkdir -p $(GEN_SRC_CPP)
	protoc --proto_path=$(PROTO_PATH) --cpp_out=$(GEN_SRC_CPP) $(SRC)

.PHONY: java
java:
	rm -rf ${GEN_SRC_JAVA}
	mkdir -p $(GEN_SRC_JAVA)
	protoc --proto_path=$(PROTO_PATH) --java_out=$(GEN_SRC_JAVA) $(SRC)

.PHONY: js
js:
	rm -rf ${GEN_SRC_JS}
	rm -rf ${GEN_SRC_TS}
	mkdir -p $(GEN_SRC_JS)
	mkdir -p $(GEN_SRC_TS)
	protoc --proto_path=$(PROTO_PATH) --js_out=import_style=commonjs,binary:$(GEN_SRC_JS) --ts_out=$(GEN_SRC_TS) $(SRC)

.PHONY: php
php:
	rm -rf ${GEN_SRC_PHP}
	mkdir -p $(GEN_SRC_PHP)
	protoc --proto_path=$(PROTO_PATH) --php_out=$(GEN_SRC_PHP) $(SRC)

.PHONY: python
python:
	rm -rf ${GEN_SRC_PYTHON}
	mkdir -p $(GEN_SRC_PYTHON)
	protoc --proto_path=$(PROTO_PATH) --python_out=$(GEN_SRC_PYTHON) $(SRC)
	@echo ""
	@echo "#####################################################################################################"
	@echo "TODO for python protobuf"
	@echo "0. Copy the generated python protobuf files to your repository"
	@echo "1. Copy 'protobuf_importpath_change.py' to root of your repository"
	@echo "2. Execute 'python protobuf_importpath_change.py -p [package_path]'. (Like iamport.protobuf_messages)"
	@echo "#####################################################################################################"
	@echo ""

.PHONY: dart
dart:
	rm -rf ${GEN_SRC_DART}
	mkdir -p $(GEN_SRC_DART)
	protoc --proto_path=$(PROTO_PATH) --dart_out=$(GEN_SRC_DART) $(SRC)

.PHONY: go
go:
	rm -rf ${GEN_SRC_GO}
	mkdir -p $(GEN_SRC_GO)
	rm -f ./go.mod
	protoc --proto_path=$(PROTO_PATH) --go_out=$(GEN_SRC_GO) --go_opt=paths=source_relative --go-grpc_out=$(GEN_SRC_GO) --go-grpc_opt=paths=source_relative --grpc-gateway_out=$(GEN_SRC_GO) --grpc-gateway_opt=paths=source_relative $(SRC)
	cp ./supplements/go/go.mod .

.PHONY: rust
rust:
	rm -rf ${GEN_SRC_RUST}
	mkdir -p $(GEN_SRC_RUST)
	protoc --proto_path=$(PROTO_PATH) --rust_out=$(GEN_SRC_RUST) $(SRC)
	cp ./supplements/rust/* $(GEN_SRC_RUST)

.PHONY: ruby
ruby:
	rm -rf ${GEN_SRC_RUBY}
	mkdir -p $(GEN_SRC_RUBY)
	protoc --proto_path=$(PROTO_PATH) --ruby_out=$(GEN_SRC_RUBY) $(SRC)

.PHONY: clean
clean:
	rm -rf $(GEN_SRC_ROOT)
	rm -f go.mod
