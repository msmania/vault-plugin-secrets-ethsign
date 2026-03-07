VGO=go # Set to vgo if building in Go 1.10
BIN_DIR=build
BINARY_NAME=ethsign
SRC_GOFILES := $(shell find . -name '*.go' -print)
.DELETE_ON_ERROR:

all: build
test: deps
	$(VGO) test  ./... -cover -coverprofile=coverage.txt -covermode=atomic
ethsign: ${SRC_GOFILES}
	# Disable CGO to run on a Alpine-based system
	CGO_ENABLED=0 $(VGO) build -o ${BIN_DIR}/${BINARY_NAME} \
		-ldflags "-X main.buildDate=`date -u +\"%Y-%m-%dT%H:%M:%SZ\"` \
		-X main.buildVersion=$(BUILD_VERSION)" \
		-tags=prod -v
build: ethsign
clean:
		$(VGO) clean
		rm -f ${BIN_DIR}/${BINARY_NAME}
deps:
		$(VGO) get
