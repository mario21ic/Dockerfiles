FROM mario21ic/nginx-php56:base
MAINTAINER Mario Inga <mario21ic@gmail.com>
ENV VERSION 1.0

RUN yum install -y \
    nginx php56-fpm

# Config Fpm
COPY ./www.conf /etc/php-fpm.d/www.conf

# Config Nginx
COPY ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

COPY ./entrypoint.sh /init/entrypoint.sh
RUN chmod 755 /init/entrypoint.sh
ENTRYPOINT ["/init/entrypoint.sh"]
