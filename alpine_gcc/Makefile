NAME=golanggcc
VERSION=1.15.2-alpine-gcc-3.12
image:
	docker build -t ${NAME}:${VERSION} .
push: image
	docker push ${NAME}:${VERSION}