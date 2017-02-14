FROM java:7-jre

MAINTAINER Mario Inga <mario21ic@gmail.com>
ENV VERSION 1.0

RUN mkdir /var/dynamodb_wd
WORKDIR /var/dynamodb_wd

RUN wget -O /var/dynamodb_wd/dynamodb_local_latest http://dynamodb-local.s3-website-us-west-2.amazonaws.com/dynamodb_local_latest
RUN tar xfz /var/dynamodb_wd/dynamodb_local_latest

ENTRYPOINT ["/usr/bin/java", "-Djava.library.path=.", "-jar", "DynamoDBLocal.jar", "-dbPath", "/var/dynamodb_local"]
CMD ["-port", "8000"]

VOLUME /var/dynamodb_local
EXPOSE 8000
