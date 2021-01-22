basename=go_jwt
version=0.0.1

image:
	echo "Building Client image"
	podman build . -t ${basename}_client:${version} --build-args MAIN_FILE=client/client.go
	podman build . -t ${basename}_server:${version} --build-args MAIN_FILE=server/server.go

comp_up:
	podman-compose -d up 

comp_down:
	podman-compose down
