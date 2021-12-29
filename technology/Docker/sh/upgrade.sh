#!/bin/sh

name=hive
port=3090
image=xxx:5000/guoguo/${name}

oldImageID=`docker images --filter=reference=${image} --format "{{.ID}}"`
docker pull $image
newImageID=`docker images --filter=reference=${image} --format "{{.ID}}"`

if [[ "${oldImageID}" != "${newImageID}" ]]; then
  docker stop $name
  docker rm $name
  sh ./run.sh
#   docker network connect nginx hive

  export CONTAINER_IP=${name}
  export CONTAINER_PORT=${port}
  envsubst '\$CONTAINER_IP \$CONTAINER_PORT' < ./conf.tpl > ~/docker/nginx/conf.d/${name}.conf
  sh ~/docker/nginx/reload.sh
else
  echo "无更新"
fi