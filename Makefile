docker=$(shell which docker)
command?=help
target?=Test

## Build library
build: clean compile test

compile:
	go build -o bin/topic-controller

test: infra _test down
_test:
	go test -count=1 -v -run $(target)

clean:
	rm -f ./bin/*

#################################
## Deploy local backing services
infra:
	$(MAKE) compose command="up -d"

logs:
	$(MAKE) compose command="logs"

down:
	$(MAKE) compose command="down"

## Runs docker compose commands
compose:
	docker run --rm --name topic-controller \
			-v $(CURDIR):/app \
			-v $(docker):/usr/bin/docker \
			-v /var/run/docker.sock:/var/run/docker.sock \
			-w /app \
			docker/compose $(command)