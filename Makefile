.PHONY: run

run:
	docker build . -t mailboxsync
	docker run \
		--mount type=bind,source=./var,target=/app/var \
		mailboxsync