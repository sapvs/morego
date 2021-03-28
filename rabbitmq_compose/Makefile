tag=localhost/gorabbit
version=0.1

image:
	podman build -t $(tag):$(version) .

pod: image
	podman pod play kube goRabbit.yaml
