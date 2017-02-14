FROM tomcat:7.0
MAINTAINER Mario Inga <mario21ic@gmail.com>

ENV VERSION 1.0

ENV TOMCAT_DIR=/usr/local/tomcat/
ENV SOLR_VERSION=4.6.0

WORKDIR /tmp
RUN wget http://archive.apache.org/dist/lucene/solr/$SOLR_VERSION/solr-$SOLR_VERSION.tgz
RUN tar -cvf /tmp/solr-$SOLR-VERSION

RUN cp -rf /tmp/solr-$SOLR-VERSION/dist/solrj-lib/*.jar $TOMCAT_DIR/lib/
COPY /tmp/solr-$SOLR-VERSION/dist/solr-$SOLR_VERSION.war $TOMCAT_DIR/webapps/aptitus2dev.war

RUN mkdir -p $TOMCAT_DIR/conf/Catalina/localhost/
COPY aptitus2dev.xml $TOMCAT_DIR/conf/Catalina/localhost/aptitus2dev.xml

WORKDIR $TOMCAT_DIR
VOLUME /data/solr/aptitus2dev/
