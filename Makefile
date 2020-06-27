docker=$(shell which docker)
command?=help
target?=Test

test: infra _test down
_test:
	go test -count=1 -v -run $(target)

infra:
	$(MAKE) compose command="up -d"

logs:
	$(MAKE) compose command="logs"

down:
	$(MAKE) compose command="down"

compose:
	docker run --rm --name topic-controller \
			-v $(CURDIR):/app \
			-v $(docker):/usr/bin/docker \
			-v /var/run/docker.sock:/var/run/docker.sock \
			-w /app \
			docker/compose $(command)