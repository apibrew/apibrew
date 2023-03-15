select sum(count)
from (select count(*) as count
      from information_schema.tables
      where table_type = 'BASE TABLE' and tables.table_schema = $1 and tables.table_name = $2
      union all
      select count (*) as count
      from information_schema.views
      where views.table_schema = $1 and views.table_name = $2) _
