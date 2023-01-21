## inside docker test
sh -c 'initdb -D /var/lib/postgresql/data' postgres
su -c 'pg_ctl start -D /var/lib/postgresql/data' postgres
su -c 'psql -h 127.0.0.1 -f /init.sql postgres' postgres

CGO_ENABLED=0 go test ./... -v