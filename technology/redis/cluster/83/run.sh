#!/bin/sh

docker run \
    -v /root/docker/redis/cluster/83/conf:/usr/local/etc/redis \
    -v /root/docker/redis/cluster/83/data:/data \
    --name redis3 \
    -p 8003:6379 \
    -d \
    redis \
    redis-server /usr/local/etc/redis/redis.conf

