#!/bin/bash
set -ex
if [ ! -f config.toml ]; then
    cp config.toml.dist config.toml
fi
go get github.com/tools/godep
godep go build

if [ $(docker ps -a | grep gs_db | wc -l) -ge 1 ]; then
    docker rm -vf gs_db
fi
if [ $(docker ps -a | grep gosense | wc -l) -ge 1 ]; then
    docker rm -vf gosense
fi
# docker run --restart=always -p 80:8080 --link db:db -v /data/www/gosense:/www --name gosense netroby/alpgo /www/gosense
docker run --restart=always -d --name gs_db  -v /mysql-data netroby/docker-mysql

while true; do
    if [ $(docker logs gs_db 2>&1 | grep "ready for connections" | wc -l)  -ge 2 ]; then
        break;
    else
        echo "not ready, waiting"
        sleep 2
    fi
done

docker cp sql/bak.sql gs_db:/root/
docker exec gs_db sh -c "mysql < /root/bak.sql"
docker run --restart=always -d -p 8080:8080 --link gs_db:db -v $(pwd):/www --name gosense netroby/alpgo /www/gosense
