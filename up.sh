#!/bin/bash
go get github.com/tools/godep
godep go build
docker rm -vf gosense
# docker run --restart=always -p 80:8080 --link db:db -v /data/www/gosense:/www --name gosense netroby/alpgo /www/gosense
docker run --restart=always -d --name db  -v /mysql-data netroby/docker-mysql
sleep 15
docker cp sql/bak.sql db:/root/
docker exec db sh -c "mysql < /root/bak.sql"
docker run --restart=always -d --link db:db -v $(pwd):/www --name gosense netroby/alpgo /www/gosense
