#!/bin/bash

ssh -tt -i blabla.pem $SSH_USER@$SERVER_IP <<EOF
    if [ ! -d $LOCAL_REPO/.git ]
    then
        git clone $REPO_SRC $LOCAL_REPO
        cd $LOCAL_REPO
    else 
        cd $LOCAL_REPO
        git checkout $TRAVIS_BRANCH
        git pull $REPO_SRC
    fi

    docker-compose stop
    echo "$TAG" > .env
    docker-compose up -d
    exit
EOF