# What is the goal?
That is a small service to run as a sidecar to secure DNS requests of your service.
You have 2 options: HTTP-over-TLS or HTTP-over-HTTPS. Check config to setup mode which you want.

# Security concerns
Certificate validation is enabled, so it more or less secure. Based on a popular 'miekg/dns' library.

# Improvements
You are welcome to add any improvements like extended logging, caching, tests, etc.

# How to $something
For build:

    make build CONTAINER_NAME=$your_name_of_the_container

For push:

    make push CONTAINER_NAME=$your_name_of_the_container

For run (container will wait requests at TCP and UDP port 10053):

    make run CONTAINER_NAME=$your_name_of_the_container

Also, you can call just `make run` to run a prebuild image from Docker Hub.
