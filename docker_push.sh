#!/bin/bash

docker build -t $IMAGE_NAME:$TAG .

echo "$DOCKER_PASSWORD" | docker login --username $DOCKER_LOGIN --password-stdin

docker tag $IMAGE_NAME:$TAG $DOCKER_LOGIN/$IMAGE_NAME:$TAG

docker push $DOCKER_LOGIN/$IMAGE_NAME:$TAG