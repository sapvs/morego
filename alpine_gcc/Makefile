NAME=golang:1.15.2-alpine3.12-gcc
VERSION=1.0
image:
	docker build -t ${NAME}:${VERSION} .
push: image
	docker push ${NAME}:${VERSION}