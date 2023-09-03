#!/bin/bash
cd "$(dirname "$0")" || exit

## inside docker test
su postgres -c 'id'
su postgres -c 'initdb -D /var/lib/postgresql/data'
su postgres -c 'pg_ctl start -D /var/lib/postgresql/data'
su postgres -c 'psql -h 127.0.0.1 -f init.sql postgres'
su postgres -c 'psql -h 127.0.0.1 -f dh_test.sql dh_test'

CGO_ENABLED=0 go test -v || exit

echo "test success"
echo "stop postgresql"
su postgres -c 'pg_ctl stop -D /var/lib/postgresql/data'