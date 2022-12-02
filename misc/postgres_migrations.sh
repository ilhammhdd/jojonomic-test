#!/bin/sh

for filename in /var/lib/postgres/migrations/*.sql; do
  echo "executing migraiton $filename"
  psql -v ON_ERROR_STOP=1 --username $POSTGRES_USER --dbname $POSTGRES_DB < $filename
done