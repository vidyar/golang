#!/bin/bash
docker rm -vf gosense
# docker run --restart=always -p 80:8080 --link db:db -v /data/www/gosense:/www --name gosense netroby/alpgo /www/gosense
docker run --restart=always -d --link db:db -v /data/www/gosense:/www --name gosense netroby/alpgo /www/gosense
