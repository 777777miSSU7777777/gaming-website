#!/bin/bash

VERSION="latest"

docker build -t $IMAGE_TAG:$VERSION .

echo "$DOCKER_PASSWORD" | docker login --username $DOCKER_LOGIN --password-stdin

docker tag $IMAGE_TAG:$VERSION $DOCKER_LOGIN/$IMAGE_TAG:$VERSION

docker push $DOCKER_LOGIN/$IMAGE_TAG:$VERSION