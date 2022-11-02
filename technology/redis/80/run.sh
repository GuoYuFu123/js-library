#!/bin/sh

docker run \
    -v /root/docker/redis/conf:/usr/local/etc/redis \
    -v /root/docker/redis/data:/data \
    --name redis \
    # -p 8001:6379 \
    -d \
    redis \
    redis-server /usr/local/etc/redis/redis.conf

