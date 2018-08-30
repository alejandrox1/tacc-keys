SSH_PORT ?= 2222
SERVER_PORT ?= 8000
CONTAINER_NAME ?= tacc-keys


.PHONY: build generate-keys run ssh clean 

build:
	docker build --force-rm --no-cache -t $(CONTAINER_NAME) .

generate-keys:
	mkdir -p ssh
	ssh-keygen -t rsa -b 4096 -C "docker" -f ssh/id_rsa -N ""
	cat ssh/id_rsa.pub >> ssh/authorized_keys

run: generate-keys build
	docker run --rm \
		-v ssh:/root/.ssh \
		-p $(SERVER_PORT):8000 -p $(SSH_PORT):22 \
		$(CONTAINER_NAME)

ssh:
	ssh -o "StrictHostKeyChecking no" -i ssh/id_rsa -p $(SSH_PORT) root@localhost

clean:
	rm -rf ssh/
