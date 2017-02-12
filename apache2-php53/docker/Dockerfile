FROM amazonlinux:latest
MAINTAINER Mario Inga <mario21ic@gmail.com>
ENV VERSION 1.0

RUN yum install -y \
    git \
    httpd \
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
    php-zip && \
    pecl install mongo && \
    curl https://getcomposer.org/composer.phar > /usr/bin/composer && \
    chmod +x /usr/bin/composer

# Config PHP
COPY ./php.ini /etc/php.ini

# Hosts
RUN su -c "echo '127.0.0.1 local.miweb.com' >> /etc/hosts"

# Application
ENV SRC_DIR=/var/www/html/
WORKDIR $SRC_DIR

# Config Apache
ENV APACHE_DIR /etc/httpd/
RUN rm -f $APACHE_DIR/conf.d/welcome.conf
COPY ./httpd.conf $APACHE_DIR/conf/httpd.conf
COPY ./local.miweb.com.conf $APACHE_DIR/conf.d/

EXPOSE 80

COPY ./entrypoint.sh /root/entrypoint.sh
ENTRYPOINT "/root/entrypoint.sh"
