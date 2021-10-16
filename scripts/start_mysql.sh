#!/bin/bash

docker run --rm -p 3306:3306 --name test-mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest

# mysql -h 127.0.0.1 -P 3306 -u root -p