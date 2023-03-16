select columns.column_name,
       columns.DATA_TYPE                                              as column_type,
       coalesce(columns.character_maximum_length, 0)                  as length,
       columns.is_nullable = 'YES'                                    as is_nullable,
       coalesce(TABLE_CONSTRAINTS.CONSTRAINT_TYPE = 'PRIMARY KEY', 0) as is_primary,
       coalesce(TABLE_CONSTRAINTS.CONSTRAINT_TYPE = 'UNIQUE', 0)      as is_unique,
       coalesce(TABLE_CONSTRAINTS.CONSTRAINT_TYPE = 'FOREIGN KEY', 0) as is_referenced,
       key_column_usage.REFERENCED_TABLE_SCHEMA                       as target_schema,
       key_column_usage.REFERENCED_TABLE_NAME                         as target_table,
       key_column_usage.REFERENCED_COLUMN_NAME                        as target_column
from information_schema.columns
         left join information_schema.key_column_usage on key_column_usage.table_name = columns.table_name and
                                                          key_column_usage.table_schema = columns.table_schema and
                                                          key_column_usage.column_name = columns.column_name
         left join information_schema.TABLE_CONSTRAINTS
                   on TABLE_CONSTRAINTS.TABLE_SCHEMA = COLUMNS.TABLE_SCHEMA and
                      TABLE_CONSTRAINTS.TABLE_NAME = COLUMNS.TABLE_NAME
where columns.table_schema = ?
  and columns.table_name = ?
order by columns.ordinal_position;
