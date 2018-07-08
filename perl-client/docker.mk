REPOSITORY := grpc-http-backdoor
TAG := perl5.28.0
IMAGE := $(REPOSITORY):$(TAG)
RUNOPT := -v $(shell pwd)/../:/work -w /work/perl-client

all: docker-build build

docker-build:
	docker build -t $(IMAGE) .

docker-shell:
	docker run --rm -it $(RUNOPT) $(IMAGE) bash

docker-clean:
	[ "$$(docker images | awk '$$1 == "$(REPOSITORY)" && $$2 == "$(TAG)" { print }' | wc -l)" != "0" ] && \
		docker rmi $(IMAGE) || true

build:
	docker run --rm $(RUNOPT) $(IMAGE) make clean all

test:
	docker run --rm $(RUNOPT) $(IMAGE) make test

clean:
	docker run --rm $(RUNOPT) $(IMAGE) make clean
	$(MAKE) -f docker.mk docker-clean

.PHONY: all docker-build docker-shell docker-clean build test
