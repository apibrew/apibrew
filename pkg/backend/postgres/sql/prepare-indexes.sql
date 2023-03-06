select i.relname as index_name,
       indisunique,
       ixs.indexdef,
       string_agg(a.attname, ',')
from pg_class t,
     pg_class i,
     pg_index ix,
     pg_indexes ixs,
     pg_attribute a
where t.oid = ix.indrelid
  and i.oid = ix.indexrelid
  and a.attrelid = t.oid
  and a.attnum = ANY (ix.indkey)
  and t.relkind = 'r'
  and ixs.indexname = i.relname
  and ixs.schemaname = $1
  and ixs.tablename = $2
  and not ix.indisprimary
group by 1, 2, 3
having count(distinct a.attname) > 1