#!/bin/sh

docker run \
    -v /root/docker/redis/cluster/84/conf:/usr/local/etc/redis \
    -v /root/docker/redis/cluster/84/data:/data \
    --name redis4 \
    -p 8004:6379 \
    -d \
    redis \
    redis-server /usr/local/etc/redis/redis.conf

