CONTAINER_NAME ?= create-keys


.PHONY: build generate-keys run ssh clean 

build:
	docker build --force-rm --no-cache -t $(CONTAINER_NAME) .

run: build
	docker run --rm -it -v ${PWD}:/go/src $(CONTAINER_NAME) /bin/bash
