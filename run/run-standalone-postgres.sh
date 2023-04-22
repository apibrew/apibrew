#!/bin/bash
cd "$(dirname "$0")" || exit

## inside docker test
chown postgres:postgres /var/lib/postgresql/data
su postgres -c 'initdb -D /var/lib/postgresql/data'
su postgres -c 'pg_ctl start -D /var/lib/postgresql/data'
su postgres -c 'psql -h 127.0.0.1 -f init.sql postgres'

pid=0

# SIGTERM-handler
term_handler() {
  if [ $pid -ne 0 ]; then
    su postgres -c 'pg_ctl stop -D /var/lib/postgresql/data'
    kill -SIGTERM "$pid"
    wait "$pid"
  fi
  exit 143; # 128 + 15 -- SIGTERM
}

trap 'kill ${!}; term_handler' SIGTERM

/bin/apibrew -log-level=debug "-init" "/app/config.json" &
pid="$!"

# wait forever
while true
do
  tail -f /dev/null & wait ${!}
done

su postgres -c 'pg_ctl stop -D /var/lib/postgresql/data'