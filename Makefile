.PHONY: install create deps

install: deps
	mkdir -p /etc/merlin && \
	cp example-config.json /etc/merlin/config.json

create:
	cd ./encoder && \
	GOOS=linux go build -o ../merlin

start:
	./merlin

development: create install start

deps:
ifeq (, $(shell which HandBrakeCLI))
	$(error "HandBrakeCLI is required but not found!")
endif