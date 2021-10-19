GO=go build
cos=unknown
short_cos=unknown
binext=.
entry=main.go

ifeq ($(OS),Windows_NT)
	cos:=windows
	short_cos:=win
	binext:=.dll
else
	unix_uname:=$(shell uname -s)
	ifeq ($(unix_uname),Darwin)
		cos:=darwin
		short_cos:=macos
		binext:=.dylib
	else
		cos:=linux
		short_cos:=linux
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

.PHONY:
	build

all: build

build:
	$(GO) -buildmode c-shared -o hson_$(short_cos)$(binext) $(entry)