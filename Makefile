CURRENT_DIR = $(shell pwd)
BUILD_DIR=$(CURRENT_DIR)/build/bin
GIT_TAG=$(shell git describe --tags --abbrev=0  2>/dev/null)
GO_BUILD_CMD_BASE=go build -ldflags "-s -w -X 'goiot/pkg/version.TagVersion=$(GIT_TAG)'" -o
OS=$(os)
OUTPUT := $(BUILD_DIR)/$(OS)
ECHO = echo
ARCH=$(arch)
##############################################################################
###   export default type define
##############################################################################
ifeq ($(ARCH), )
ARCH :=x86
endif
export ARCH

#############################################################################
###   export default os define
#############################################################################
ifeq ($(OS), )
OS :=osx
endif
export OS
#############################################################################
###   other platform
#############################################################################
ifeq ($(os), amd64)
	GOOS=linux
    GOARCH=amd64
    OUTPUT := $(BUILD_DIR)/$(GOOS)/$(os)
    GO_BUILD_CMD=CGO_ENABLED=0 $(GO_BUILD_CMD_BASE)
    CC=x86_64-linux-gnu-gcc
    GCC=x86_64-linux-gnu-gcc
    CXX=x86_64-linux-gnu-g++
    AR=x86_64-linux-gnu-ar
    LINK=x86_64-linux-gnu-gcc
    CGO_CFLAGS=
    CGO_LDFLAGS=
	export CGO_LDFLAGS LD_LIBRARY_PATH CGO_CFLAGS CC GCC CXX AR LINK GOOS GOARCH

else ifeq ($(os), arm32)
	GOOS=linux
    GOARCH=arm
    GOARM=7
    OUTPUT := $(BUILD_DIR)/$(GOOS)/$(os)
    GO_BUILD_CMD=CGO_ENABLED=1 $(GO_BUILD_CMD_BASE)
    CC=arm-linux-gnueabihf-gcc
    GCC=arm-linux-gnueabihf-gcc
    CXX=arm-linux-gnueabihf-g++
    AR=arm-linux-gnueabihf-ar
    LINK=arm-linux-gnueabihf-gcc
    CGO_CFLAGS =
    CGO_LDFLAGS =
	export CGO_LDFLAGS LD_LIBRARY_PATH CGO_CFLAGS CC GCC CXX AR LINK GOOS GOARCH GOARM

else ifeq ($(os), arm64)
	GOOS=linux
    GOARCH=arm64
    GOARM=7
    OUTPUT := $(BUILD_DIR)/$(GOOS)/$(os)
    GO_BUILD_CMD=CGO_ENABLED=1 $(GO_BUILD_CMD_BASE)
    CC=aarch64-linux-gnu-gcc
    GCC=aarch64-linux-gnu-gcc
    CXX=aarch64-linux-gnu-g++
    AR=aarch64-linux-gnu-ar
    LINK=aaarch64-linux-gnu-gcc
    CGO_CFLAGS=
    CGO_LDFLAGS=
	export CGO_LDFLAGS LD_LIBRARY_PATH CGO_CFLAGS CC GCC CXX AR LINK GOOS GOARCH GOARM

else ifeq ($(os), osx)
	GOOS=darwin
	GOARCH=amd64
	OUTPUT := $(BUILD_DIR)/$(GOOS)/$(os)
	export GOOS GOARCH GO_BUILD_FLAG

else
endif


all: date goiot
goiot: date goiot
date:
	@$(ECHO) "package version \n\nconst (Date = \"`date +%y%m%d.%H.%M.%S`\")" > ./pkg/version/versionDate.go
goiot:
	cc=$(CC) $(GO_BUILD_CMD) $(OUTPUT)/$@-api $(GO_BUILD_FLAG) ./apps/iotcore/api/*.go
	cc=$(CC) $(GO_BUILD_CMD) $(OUTPUT)/$@-rpc $(GO_BUILD_FLAG) ./apps/iotcore/rpc/*.go