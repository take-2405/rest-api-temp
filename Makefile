GOCMD=go
DOCKERCMD=docker
GO_RUN=$(GOCMD) run
GO_BUILD=$(GOCMD) build
DOCKER_BUILD=$(DOCKERCMD) build
DOCKER_RUN=$(DOCKERCMD) run

all:
clean:
	rm server
build:
	$(GO_BUILD) -o hello main.go
run:
	$(GO_RUN) main.go
docker-build:
	$(DOCKER_BUILD) ./ -t take-2405/rest-api-temp:0.1.0
docker-run:
	$(DOCKER_RUN) -d -p 8080:8080 take-2405/rest-api-temp:0.1.0
compose-build:
	docker-compose -f build/docker-compose.yml up -d --build
compose-clean:
	docker-compose -f build/docker-compose.yml down --rmi all --volumes