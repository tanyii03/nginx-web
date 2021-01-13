#!/bin/sh
envsubst '$API_SERVER' < /etc/nginx/nginx.conf.tpl > /etc/nginx/nginx.conf
env nginx
chmod + nginx-web
./nginx-web
