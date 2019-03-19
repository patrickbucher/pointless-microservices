#!/bin/sh

source ./services.list

echo -n "$SERVICES" | while read -r LINE; do
    NAME=$(echo $LINE | cut -d':' -f1)
    PORT=$(echo $LINE | cut -d':' -f2)
    docker build . --build-arg PORT=$PORT --build-arg FILE_BASE_NAME=$NAME -t $NAME
    docker rm -f $NAME
    docker run -dit -p $PORT:$PORT --name $NAME $NAME
done
