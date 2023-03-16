select count(*)
from (select table_schema, table_name
      from information_schema.tables
      where table_type = 'BASE TABLE'
      union all
      select table_schema, table_name
      from information_schema.views) _
where table_schema = ?
  and table_name = ?
