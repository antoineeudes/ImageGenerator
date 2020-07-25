# ImageGenerator

This software is powered by Go. It is a demonstration of image generation and treatment.

# Installation
## Prerequisites
- [Docker](https://docs.docker.com/get-docker/)

## Building the image
```bash
cd /path/to/this/directory
docker build --tag go-image-generator .
```

# Usage
```bash
docker run -ti -v /path/to/this/directory:/go/src/ImageGenerator go-image-generator
go run main.go
```

The image generated will appear on your host machine at `/path/to/this/directory`.