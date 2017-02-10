FROM mario21ic/nginx-php56:base
MAINTAINER Mario Inga <mario21ic@gmail.com>
ENV VERSION 1.0

RUN yum install -y git

# Composer
ENV COMPOSER_HOME=$APP_DIR/.composer
RUN curl https://getcomposer.org/composer.phar > /usr/bin/composer && chmod +x /usr/bin/composer

# Entrypoint
COPY ./entrypoint.sh /root/entrypoint.sh
RUN chmod +x /root/entrypoint.sh
ENTRYPOINT ["/root/entrypoint.sh"]
