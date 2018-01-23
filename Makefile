.PHONY: clean deps build run test

clean:
	@rm -fR vendor

deps:
	dep ensure

update:
	dep ensure -update

build: deps
	go build

run: 
	go run main.go this is args

test: 
	go test
