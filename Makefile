# Go build command
GO=go build
# Current OS, default mark as unknown
cos=unknown
# Extension of dynamic library
binext=.
# Entry point of the program
entry=main.go

ifeq ($(OS),Windows_NT)
	cos:=windows
	binext:=.dll
else
	unix_uname:=$(shell uname -s)
	ifeq ($(unix_uname),Darwin)
		cos:=darwin
		binext:=.dylib
	else
		cos:=linux
		binext:=.so
	endif
endif

ifndef CGO_ENABLED
	$(error CGO_ENABLED is undefined)
else
	ifeq ($(CGO_ENABLED),0)
		$(error Please set CGO_ENABLED as 1)
	endif
endif

ifndef GOOS
	export GOOS=$(cos)
endif

.PHONY:
	build

all: build clean

build:
	$(GO) -buildmode c-shared -o hashjson$(binext) $(entry)