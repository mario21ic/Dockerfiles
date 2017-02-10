FROM mario21ic/apach2-php53:latest

RUN yum install -y httpd

# Config Apache
ENV APACHE_DIR /etc/httpd/
RUN rm -f $APACHE_DIR/conf.d/welcome.conf
COPY ./httpd.conf $APACHE_DIR/conf/httpd.conf
COPY ./local.miweb.com.conf $APACHE_DIR/conf.d/

EXPOSE 80

COPY ./init.sh /init/init.sh
RUN chmod +x /init/init.sh
ENTRYPOINT "/init/init.sh"
