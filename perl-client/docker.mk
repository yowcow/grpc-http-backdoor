IMAGE := grpc-http-backdoor:perl5.28.0
RUNOPT := -v $(shell pwd)/../:/work -w /work/perl-client

all: docker-build build

docker-build:
	docker build -t $(IMAGE) .

docker-shell:
	docker run --rm -it $(RUNOPT) $(IMAGE) bash

build:
	docker run --rm $(RUNOPT) $(IMAGE) make clean all

test:
	docker run --rm $(RUNOPT) $(IMAGE) make test

.PHONY: all docker-build docker-shell build test
