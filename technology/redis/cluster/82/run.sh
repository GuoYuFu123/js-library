#!/bin/sh

docker run \
    -v /root/docker/redis/cluster/82/conf:/usr/local/etc/redis \
    -v /root/docker/redis/cluster/82/data:/data \
    --name redis2 \
    -p 8002:6379 \
    -d \
    redis \
    redis-server /usr/local/etc/redis/redis.conf

