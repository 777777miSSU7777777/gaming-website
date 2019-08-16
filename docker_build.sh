#!/bin/bash

VERSION="latest"

docker build -t $IMAGE_TAG:$VERSION .
