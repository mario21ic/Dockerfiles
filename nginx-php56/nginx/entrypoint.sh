#!/bin/bash

/usr/sbin/php-fpm -D && /usr/sbin/nginx -g 'daemon off;' && tail -f /var/log/nginx/local.miweb.com.error.log
