FROM java:8

MAINTAINER Mario Inga <mario21ic@gmail.com>
ENV VERSION 1.0

RUN mkdir /var/dynamodb_wd
WORKDIR /var/dynamodb_wd

RUN wget -O /var/dynamodb_wd/dynamodb_local_latest https://s3-us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.tar.gz
RUN tar xfz /var/dynamodb_wd/dynamodb_local_latest

ENTRYPOINT ["/usr/bin/java", "-Djava.library.path=.", "-jar", "DynamoDBLocal.jar", "-dbPath", "/var/dynamodb_local"]
CMD ["-port", "8000"]

VOLUME /var/dynamodb_local
EXPOSE 8000
