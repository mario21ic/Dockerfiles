FROM php:5.4-apache
MAINTAINER Mario Inga <mario21ic@gmail.com>


RUN apt-get update && apt-get install -y \
    git mysql-client php5-dev php5-mysql php-pear curl \
    libssl-dev libmemcached-dev zlib1g-dev \
    libncurses5-dev libxml2-dev gcc libpcre3-dev \
    mbstring mysql mysqli pdo pdo_mysql mongodb


RUN curl -L http://pecl.php.net/get/mongo-1.6.9.tgz >> /usr/src/php/ext/mongo-1.6.9.tgz && \
    tar -xvf /usr/src/php/ext/mongo-1.6.9.tgz -C /usr/src/php/ext/ && \
    rm /usr/src/php/ext/mongo-1.6.9.tgz && \
    docker-php-ext-install mongo-1.6.9

# Phalcon
RUN git clone --depth=1 -b 1.3.4 git://github.com/phalcon/cphalcon.git && cd cphalcon/build && ./install

# Config PHP
#COPY ./php.ini /usr/local/etc/php/

# Config Apache
ENV APACHE_DIR /etc/apache2
RUN rm -f $APACHE_DIR/sites-enabled/*
COPY ./local.miweb.com.conf $APACHE_DIR/sites-available/
RUN a2ensite local.miweb.com
RUN a2enmod rewrite
