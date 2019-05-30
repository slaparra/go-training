# Docker
https://www.docker.com/get-started

Golang official docker image
- https://hub.docker.com/_/golang

Todd McLeod documentation and some examples:
- https://github.com/GoesToEleven/golang-web-dev/tree/master/043_docker

# Docker with Go and RabbitMq
Example:
- See: [docker-compose.yml](../src/16-docker/docker-compose.yml) and the [example configuration](../src/16-docker/)
- Go to path `src/16-docker` and execute:
    - `docker-compose up`
- Visit:
    - [http://localhost:8080](http://localhost:8080) to send a message
    - [http://localhost:15672](http://localhost:15672) is the RabbitMq admin (admin/1234)

## Other links
- [RabbitMq get started](https://www.rabbitmq.com/getstarted.html)
- [Docker dev environment with Go Modules and live code reloading](https://threedots.tech/post/go-docker-dev-environment-with-go-modules-and-live-code-reloading/)

