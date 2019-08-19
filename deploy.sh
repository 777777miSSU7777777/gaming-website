#!/bin/bash

ssh -tt -i blabla.pem $SSH_USER@$SERVER_IP <<EOF
    if [ ! -d $LOCAL_REPO/.git ]
    then
        git clone $REPO_SRC $LOCAL_REPO
        cd $LOCAL_REPO
    else 
        cd $LOCAL_REPO
        git checkout $TRAVIS_BRANCH
        git pull
    fi

    docker-compose stop
    echo "TAG=$(git rev-parse HEAD)" > .env
    docker-compose up -d
    exit
EOF