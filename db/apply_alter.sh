#!/bin/bash

DB_DIR=$(cd $(dirname $0) && pwd)
cd $DB_DIR

mysql -uroot isubata < ./alter.sql
