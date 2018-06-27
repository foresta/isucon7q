#/bin/bash

on() {
    mysql -uroot -e "set global slow_query_log_file = '/var/log/mysql/mysql-slow.log';"
    mysql -uroot -e "set global long_query_time = 0;"
    mysql -uroot -e "set global slow_query_log = ON;"
}

off() {
     mysql -uroot -e "set global long_query_time = 10;"
     mysql -uroot -e "set global slow_query_log = OFF;"
}

if [ "$1" = "off" ]; then
    off
else
    on
fi
