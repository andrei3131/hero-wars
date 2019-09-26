[![Build Status](https://travis-ci.org/andrei3131/hero-wars.svg?branch=master)](https://travis-ci.org/andrei3131/hero-wars)

# Running Hero Wars

## 1. Build & run locally

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

```
docker pull andrei3131/hero-wars
docker run -it hero-wars
```


