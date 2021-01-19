image:
	podman build . -t sudosapan/goactions:0.1

test:
	@go test */**

build:
	go build 

run: image
	podman run -d --name goaction  -p 8080:8080 sudosapan/goactions:0.1

stop:
	podman stop goaction
	podman rm goaction

