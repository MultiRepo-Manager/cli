#!/bin/bash

IMAGE=docker.pkg.github.com/multirepo-manager/backend/devenv
TAG=develop

if [ $# -lt 1 ]
then
  echo "Usage: devenv.sh PATH <config-repo>"
  exit 0
fi

if [ $# -eq 2 ]
then
  git clone $2 $1/.devenv
fi

if [ ! -d $1/.devenv ]
then
  mkdir -p $1/.devenv
  echo "{\"name\":\"New project\",\"repos\":[]}" > $1/.devenv/config.json
fi

docker run -p 127.0.0.1:8080:8080/tcp -v $1:/app/src/workspace -v $HOME/.ssh:/root/.ssh $IMAGE:$TAG

