.PHONY: build
build:
	@echo "Building..."
	# TODO build client and make sure binary is available
	# TODO env variables for tags and repo name
	docker build . -t natedroid/go-orchestra


.PHONY: push
push:
	docker image push natedroid/go-orchestra

.PHONY: test-all
	go test -v ./...