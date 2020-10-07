#!/bin/sh

PORT=8880
IMAGE_NAME=go-voting-app
IMAGE_TAG=0.0.2
CONTAINER_NAME=go-running-app

# Build docker image
sudo docker build -t $IMAGE_NAME:$IMAGE_TAG .

# Run container
sudo docker run -it --rm -d -p $PORT:$PORT --name $CONTAINER_NAME $IMAGE_NAME:$IMAGE_TAG
