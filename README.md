# docker-hello-world

> Simple web application that renders the hostname of the container where it's being run

# Quickstart

Simply run in the terminal
```
$ docker run --name hello-world -d -p 8080:80 containersol/hello-world
```

Then you can access the application on [http://localhost:8080]()

# Local Development
## Requirements
  - Docker
  - Make

## Building

```
$ make build
```
