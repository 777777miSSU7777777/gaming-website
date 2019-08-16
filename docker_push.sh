#!/bin/bash

echo "$DOCKER_PASSWORD" | docker login --username $DOCKER_LOGIN --password-stdin

docker tag $IMAGE_TAG:$VERSION_TAG $DOCKER_LOGIN/$IMAGE_TAG:$VERSION_TAG

docker push $DOCKER_LOGIN/$IMAGE_TAG:$VERSION_TAG