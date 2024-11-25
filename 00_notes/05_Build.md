# Build

## Docker

run script in scripts/buildprod.sh to build the binary

```
./scripts/buildprod.sh
```

build the Docker image locally (replace the namespace with Docker Hub username)

```
docker build -t DOCKERHUB_NAMESPACE/notely:latest .
```

run the Docker image locally

```
docker run -e PORT=8080 -p 8080:8080 DOCKERHUB_NAMESPACE/notely:latest
```

## Continuous Deployment (CD)
