CONTAINER_NAME="onlinehead/dns-to-dns-tls"
build:
	docker build -t $(CONTAINER_NAME) .
push:
	docker push $(CONTAINER_NAME)
run:
	docker run -p 10053:53 -p 10053:53/udp $(CONTAINER_NAME)
