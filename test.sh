#!/bin/sh

source ./services.list

echo -n "$SERVICES" | while read -r LINE; do
    NAME=$(echo $LINE | cut -d':' -f1)
    PORT=$(echo $LINE | cut -d':' -f2)
    curl localhost:$PORT/$NAME
done
