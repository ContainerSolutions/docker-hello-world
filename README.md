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

## Pushing the image to DockerHub
1. Make sure that you have the following env variables set before running these steps:
   - `DOCKER_USER`
   - `DOCKER_PASS`

2. Authenticate with DockerHub:
   ```
   $ make login
   ```

3. Push the image and tags:
   ```
   $ make push
   ```
