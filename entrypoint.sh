#!/bin/sh

set -x

cd /app || exit

if [ ! -f "/app/expose/install.lock" ];then
  # busuanzi js API address
  if [ -n "$API_SERVER" ];then
    sed -i "s|http://127.0.0.1:8080/api|$API_SERVER|g" dist/busuanzi.js

    echo "Replace API_SERVER to $API_SERVER"
  fi

  mv dist /app/expose/dist
  mv config.yaml /app/expose/config.yaml
  touch /app/expose/install.lock
fi

exec /app/busuanzi -c /app/expose/config.yaml -d /app/expose/dist