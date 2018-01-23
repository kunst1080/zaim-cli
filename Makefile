NAME := zaim
SRCS := $(shell find . -type f -name '*.go')
ARG   = verify

bin/${NAME}: deps $(SRCS)
	go build -o bin/$(NAME)

.PHONY: clean deps build run test

clean:
	rm -rf bin/*
	rm -fR vendor

deps:
	dep ensure

update:
	dep ensure -update

run-jq: 
	go run *.go ${ARG} | jq .

run: 
	go run *.go ${ARG}
