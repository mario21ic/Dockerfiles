#!/bin/bash

docker-compose exec --user=$(id -u) backend composer $*
