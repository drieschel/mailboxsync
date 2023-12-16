.PHONY: run

CONCURRENT_SYNCS ?= 3

run:
	docker build . -t mailboxsync
	docker run \
		--mount type=bind,source=./var,target=/app/var \
		mailboxsync --concurrent-syncs $(CONCURRENT_SYNCS)