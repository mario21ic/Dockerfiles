FROM amazonlinux:latest
MAINTAINER Mario Inga <mario21ic@gmail.com>
ENV VERSION 1.0

RUN yum install -y \
    gcc openssl-devel \
    php56-devel \
    php56-curl php56-gd \
    php56-mbstring \
    php56-mysqlnd \
    php56-mcrypt \
    php56-pecl-memcache \
    php56-pecl-memcached \
    php56-pecl-apc \
    php56-pecl-imagick \
    php56-soap

RUN pecl install mongodb
RUN pecl install xdebug

# Config PHP
COPY ./php.ini /etc/php.ini
COPY ./51-mongodb.ini /etc/php-5.6.d/51-mongodb.ini
COPY ./52-xdebug.ini /etc/php-5.6.d/52-xdebug.ini

# Hosts
#RUN su -c "echo '127.0.0.1 local.miweb.com' >> /etc/hosts"

ENV APP_DIR=/app
WORKDIR $APP_DIR
