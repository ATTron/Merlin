GOCMD=go
GOGET=$(GOCMD) get
UNAME:=$(shell uname -s)
NUX:= $(shell command ls /etc/yum.repos.d/*nux* 2> /dev/null)
HANDBRAKE:=$(shell command HandBrakeCLI --help 2> /dev/null)
REDHATOS:=$(shell command cat /etc/redhat-release 2> /dev/null)

.PHONY: install create deps development-deps install-handbrake

install: deps
	mkdir -p /etc/merlin && \
	cp example-config.json /etc/merlin/config.json

create:
	cd ./encoder && \
	GOOS=linux go build -o ../merlin

start:
	./merlin

development: development-deps create install start

development-deps: deps
	go get github.com/rjeczalik/notify

deps:
ifeq (, $(shell which HandBrakeCLI))
	$(error "HandBrakeCLI is required but not found!")
endif

install-handbrake:
ifeq ($(UNAME), Linux)
ifdef REDHATOS
ifndef $(NUX)
	rpm -qf /etc/yum.repos.d/*nux*
endif
ifndef $(HANDBRAKE)
	@sudo yum -y install http://li.nux.ro/download/nux/dextop/el7/x86_64/nux-dextop-release-0-5.el7.nux.noarch.rpm
	@sudo yum -y install HandBrake-cli
else
	$(shell echo "HandBrakeCLI has already been installed")
endif
endif
endif