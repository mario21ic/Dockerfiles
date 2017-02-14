FROM amazonlinux:latest
MAINTAINER Mario Inga <mario21ic@gmail.com>
ENV VERSION 1.0

ENV APP_DIR=/app
ENV COMPOSER_HOME=$APP_DIR/.composer

RUN yum install -y \
    git gcc openssl-devel \
    nginx php56-fpm \
    php56-devel \
    php56-curl php56-gd \
    php56-mbstring \
    php56-mysqlnd \
    php56-mcrypt \
    php56-pecl-memcache \
    php56-pecl-memcached \
    php56-pecl-apc \
    php56-pecl-imagick \
    php56-soap && \
    pecl install mongodb && \
    pecl install xdebug && \
    curl https://getcomposer.org/composer.phar > /usr/bin/composer && \
    chmod +x /usr/bin/composer

# Config PHP
COPY ./php.ini /etc/php.ini
COPY ./51-mongodb.ini /etc/php-5.6.d/51-mongodb.ini
COPY ./52-xdebug.ini /etc/php-5.6.d/52-xdebug.ini

# Config Fpm
COPY ./www.conf /etc/php-fpm.d/www.conf

# Config Nginx
COPY ./nginx.conf /etc/nginx/nginx.conf

# Hosts
RUN su -c "echo '127.0.0.1 local.miweb.com' >> /etc/hosts"

WORKDIR $APP_DIR

EXPOSE 80

COPY ./entrypoint.sh /init/entrypoint.sh
ENTRYPOINT ["/init/entrypoint.sh"]
