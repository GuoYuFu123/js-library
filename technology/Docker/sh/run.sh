#!/bin/sh

docker run --name hive \
        --restart=always \
        --network=mysql \
        -d xxxx:5000/guoguo/docker