FROM ruby:2.1
MAINTAINER Mario Inga <mario21ic@gmail.com>
ENV VERSION 1.0

RUN gem install fakes3 -v 0.2.4

RUN mkdir -p /data/fakes3

ENTRYPOINT ["/usr/local/bundle/bin/fakes3"]
CMD ["-r",  "/data/fakes3", "-p",  "4567"]

VOLUME /data/fakes3
EXPOSE 4567
