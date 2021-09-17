build:
	DOCKER_BUILDKIT=1
	docker-compose -f docker-compose.yaml build

up:
	docker-compose up

update-core:
	cd conductor/
	go get -v -u github.com/nate-droid/core
	cd ..
	cd musician
	go get -v -u github.com/nate-droid/core
	cd ..