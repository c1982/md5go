# md5go

Example for my blog post

## build

docker build --build-arg VERSION=v1.0.0 -t md5go:latest .

## run

docker run -d --name md5go -p 8080:8080 md5go:latest