#!/bin/bash
go get github.com/tools/godep
godep go build
docker rm -vf gs_db
docker rm -vf gosense
# docker run --restart=always -p 80:8080 --link db:db -v /data/www/gosense:/www --name gosense netroby/alpgo /www/gosense
docker run --restart=always -d --name gs_db  -v /mysql-data netroby/docker-mysql
sleep 15
docker cp sql/bak.sql gs_db:/root/
docker exec gs_db sh -c "mysql < /root/bak.sql"
docker run --restart=always -d --link gs_db:db -v $(pwd):/www --name gosense netroby/alpgo /www/gosense
