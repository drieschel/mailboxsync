.PHONY: start stop

APP_NAME ?= mailboxsync

DOCKER_COMPOSE := docker-compose -p $(APP_NAME)

start: stop
	$(DOCKER_COMPOSE) build --pull
	$(DOCKER_COMPOSE) up --force-recreate
#	$(DOCKER_COMPOSE) up -d --force-recreate

stop:
	$(DOCKER_COMPOSE) stop
	$(DOCKER_COMPOSE) rm -f
