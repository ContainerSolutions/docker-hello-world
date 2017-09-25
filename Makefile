IMG := containersol/hello-world
TAG := $$(git log -1 --pretty=%H)
NAME := ${IMG}:${TAG}

run:
	@docker run --rm -it -p 4321:80 ${NAME}

build:
	@docker build -t ${NAME} .
	@docker tag ${NAME} ${IMG}:latest

