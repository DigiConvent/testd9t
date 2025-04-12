drop view if exists _permission_check;
create view _permission_check as 
with recursive hierarchy(value, str, accumulated) as (
   select
      '',
      name || '.',
      ''
   from permissions
   union all
   select
      substr(str, 1, instr(str, '.') - 1),
      substr(str, instr(str, '.') + 1),
      accumulated || case when accumulated = '' then '' else '.' end || substr(str, 1, instr(str, '.') - 1)
   from hierarchy
   where str != ''
)
select 
   distinct(accumulated), 
   (select count(*) from permissions where name = accumulated) as "exists" from hierarchy where value != '' and "exists" = 0;
select * from _permission_check;

create trigger if not exists after_insert_permission
after insert on permissions
for each row
begin
   insert into permissions (name, 'meta') select accumulated, '->after_insert_permission:' || accumulated from _permission_check;
end;