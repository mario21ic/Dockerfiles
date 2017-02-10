FROM amazonlinux:latest
MAINTAINER Mario Inga <mario21ic@gmail.com>
ENV VERSION 1.0

RUN yum install -y \
    php \
    gcc openssl-devel \
    php-pear php-devel \
    php-pecl-apc \
    php-bcmath \
    php-curl \
    php-dom \
    php-gd \
    php-iconv \
    php-intl \
    php-mcrypt \
    php-mysqlnd \
    php-mbstring \
    php-pdo_mysql \
    php-pdo \
    php-phar \
    php-pecl-memcache \
    php-pecl-memcached \
    php-soap \
    php-xml \
    php-zip

RUN pecl install mongo

# Config PHP
COPY ./php.ini /etc/php.ini

# Hosts
RUN su -c "echo '127.0.0.1 local.miweb.com' >> /etc/hosts"

# Application
ENV SRC_DIR=/var/www/html/
WORKDIR $SRC_DIR
