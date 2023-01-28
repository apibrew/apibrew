## inside docker test
su postgres -c 'id'
su postgres -c 'initdb -D /var/lib/postgresql/data'
su postgres -c 'pg_ctl start -D /var/lib/postgresql/data'
su postgres -c 'psql -h 127.0.0.1 -f test/init.sql postgres'
su postgres -c 'psql -h 127.0.0.1 -f test/dh_test.sql dh_test'

cd test || exit
CGO_ENABLED=0 go test

su postgres -c 'pg_ctl stop -D /var/lib/postgresql/data'