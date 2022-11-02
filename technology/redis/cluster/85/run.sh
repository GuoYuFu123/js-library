#!/bin/sh

docker run \
    -v /root/docker/redis/cluster/85/conf:/usr/local/etc/redis \
    -v /root/docker/redis/cluster/85/data:/data \
    --name redis5 \
    -p 8005:6379 \
    -d \
    redis \
    redis-server /usr/local/etc/redis/redis.conf

