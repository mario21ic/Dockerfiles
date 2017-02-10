FROM amazonlinux:latest
MAINTAINER Mario Inga <mario21ic@gmail.com>

RUN yum install -y gcc-c++ postgresql93-devel libxml2-devel \
        ruby23 ruby23-devel rubygems23 git postgresql93-libs \
        nodejs cmake gcc \
        && gem install bundler \
        && rpm -ivh http://dl.fedoraproject.org/pub/epel/6/x86_64/epel-release-6-8.noarch.rpm \
        && yum --enablerepo=epel install -y nodejs

RUN mkdir -p /apps/web
WORKDIR /apps/web
COPY Gemfile /apps/web/Gemfile
COPY Gemfile.lock /apps/web/Gemfile.lock
RUN bundle install
