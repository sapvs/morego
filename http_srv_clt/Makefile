podname=clientserverpod
servertag=server
clienttag=client
version=0.1

clientimage:
	podman build -f client -t $(clienttag):$(version)

serverimage:
	podman build -f server -t $(servertag):$(version)

image: clientimage serverimage

runserver: serverimage
	podman run -p 8080:8080 -h server $(servertag):$(version) 

runclient: clientimage
	podman run -p 8081:8081 -h client $(clienttag):$(version)

pod: image
	podman pod rm -f client_server
	podman play kube srv_clt.yaml