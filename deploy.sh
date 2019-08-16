#!/bin/bash
ssh -tt -i blabla.pem $SSH_USER@$SERVER_IP <<EOF
    if [ ! -d $LOCAL_REPO/.git ]
    then
        git clone $REPO_SRC $LOCAL_REPO
        cd $LOCAL_REPO
    else 
        cd $LOCAL_REPO
        git pull $REPO_SRC
    fi

    git checkout deploy
    docker-compose down
    docker-compose up -d
EOF