#!/bin/bash

IMAGE=$DOCKER_LOGIN:$IMAGE_NAME

docker build -t $IMAGE:$TAG .

echo "$DOCKER_PASSWORD" | docker login --username $DOCKER_LOGIN --password-stdin

docker tag $IMAGE:$TAG $DOCKER_LOGIN/$IMAGE:$TAG

docker push $DOCKER_LOGIN/$IMAGE:$TAG