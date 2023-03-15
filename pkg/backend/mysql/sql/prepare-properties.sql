select columns.column_name,
       columns.udt_name                                         as column_type,
       columns.character_maximum_length                         as length,
       columns.is_nullable = 'YES'                              as is_nullable,
       exists(SELECT 1
              FROM pg_constraint c
              WHERE contype = 'p'
                and conname = key_column_usage.constraint_name) as is_primary,
       exists(SELECT 1
              FROM pg_constraint c
              WHERE contype = 'u'
                and conname = key_column_usage.constraint_name) as is_unique,
       column_fkey.constraint_def is not null                   as is_referenced,
       column_fkey.target_schema,
       column_fkey.target_table,
       column_fkey.target_column
from information_schema.columns
         left join information_schema.key_column_usage on key_column_usage.table_name = columns.table_name and
                                                          key_column_usage.table_schema = columns.table_schema and
                                                          key_column_usage.column_name = columns.column_name
         left join LATERAL (SELECT nspname                                                       as table_schema,
                           conname,
                           contype,
                           pg_get_constraintdef(c.oid)                                   as constraint_def,
                           (SELECT nspname FROM pg_namespace WHERE oid = f.relnamespace) AS target_schema,
                           f.relname                                                     AS target_table,
                           (SELECT a.attname
                            FROM pg_attribute a
                            WHERE a.attrelid = f.oid
                              AND a.attnum = c.confkey[1]
                              AND a.attisdropped = false)                                AS target_column
                    FROM pg_constraint c
                             JOIN pg_namespace n ON n.oid = c.connamespace
                             LEFT JOIN pg_class f ON f.oid = c.confrelid
                             LEFT JOIN pg_class m ON m.oid = c.conrelid
                    WHERE contype = 'f' and conname = key_column_usage.constraint_name
                      and c.conrelid IN (SELECT oid FROM pg_class c WHERE c.relkind = 'r')) column_fkey
                   on true
where columns.table_schema = $1
  and columns.table_name = $2
order by columns.ordinal_position