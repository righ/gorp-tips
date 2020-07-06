#!/bin/bash

set -e

host="$1"
shift
cmd="$@"

until mysql -uusr -ppw -hmysql db -e 'show databases' 2>/dev/null; do
  sleep 1
done

>&2 echo "MySQL is up - executing command"
exec $cmd