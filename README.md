[![Build Status](https://travis-ci.org/andrei3131/hero-wars.svg?branch=master)](https://travis-ci.org/andrei3131/hero-wars)

# Running Hero Wars

Player configurations are specified in config.yml.

## 1. Build & run locally in an environment where Go is installed

```
cd hero-wars
make build
make test (This is optional)
./bin/hero-wars
```

## 2. Build & run locally using Docker

### Build Docker image

```
docker build -t hero-wars -f Dockerfile .
```

### Run container

This will run a simulation of the game during the lifetime of this container.

```
docker run -it hero-wars
```

## 3. Pull already built image from Docker Hub and run the simulation


The commands below directly run the simulation in a Docker container.
```
docker pull andrei3131/hero-wars:latest
docker run -it hero-wars
```

The commands below allow the user to modify the configuration file of the heros (config.yml), by providing an interactive console to the container.
```
docker pull andrei3131/hero-wars:configurable
docker run -it hero-wars:configurable
```



