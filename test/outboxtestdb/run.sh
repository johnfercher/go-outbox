#!/bin/bash

docker run --volume=./data:/var/lib/mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=supersecret -e MYSQL_DATABASE=OutboxTestDb -e MYSQL_USER=AdminUser -e MYSQL_PASSWORD=AdminPassword outbox-test-db