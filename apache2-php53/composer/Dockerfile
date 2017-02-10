FROM mario21ice/apache2-php53:latest
MAINTAINER Mario Inga <mario21ic@gmail.com>
ENV VERSION 1.0

RUN yum install -y git

# Composer
RUN curl https://getcomposer.org/composer.phar > /usr/bin/composer
RUN chmod +x /usr/bin/composer

# Entrypoint
COPY ./entrypoint.sh /root/entrypoint.sh
RUN chmod +x /root/entrypoint.sh
ENTRYPOINT ["/root/entrypoint.sh"]
