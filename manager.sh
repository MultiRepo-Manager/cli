#!/bin/bash

IMAGE=docker.pkg.github.com/multirepo-manager/backend/manager
TAG=develop

if [ $# -lt 1 ]
then
  echo "Usage: manager.sh PATH <config-repo>"
  exit 0
fi

if [ $# -eq 2 ]
then
  git clone $2 $1/.manager
fi

if [ ! -d $1/.manager ]
then
  mkdir -p $1/.manager
  echo "{\"name\":\"New project\",\"repos\":[]}" > $1/.manager/config.json
fi

docker run -p 127.0.0.1:8080:8080/tcp -v $1:/app/src/workspace -v $HOME/.ssh:/root/.ssh $IMAGE:$TAG

