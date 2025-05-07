psql -f init.sql postgres
psql -c 'drop schema public cascade;' dh_test;
psql -c 'drop schema test1 cascade;' dh_test;
psql -c 'drop schema test2 cascade;' dh_test;
psql -c 'create schema public;' dh_test;
psql -f dh_test.sql dh_test;