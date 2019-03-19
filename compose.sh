#!/bin/sh

SERVICES="date:3112
time:2359
dice:6666
onlinecheck:1234
"

echo -n "$SERVICES" | while read -r LINE; do
    NAME=$(echo $LINE | cut -d':' -f1)
    PORT=$(echo $LINE | cut -d':' -f2)
    docker build . --build-arg PORT=$PORT --build-arg FILE_BASE_NAME=$NAME -t $NAME && \
    docker run -dit -p $PORT:$PORT $NAME
done
