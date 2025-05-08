create trigger check_status_overlap_insert
before insert on permission_group_has_user
for each row
-- check if the permission_group if of type status
when exists (
   select 1 from permission_groups 
   where id = new.permission_group and meta = 'status'
)
begin
   select raise(abort, 'overlapping time span for status assignment is not allowed')
   where exists (
      select 1 
      from permission_group_has_user
      where permission_group = new.permission_group
         and user = new.user
         and (
            (start < new.start and new.start < "end")
            or (start < new."end" and new."end" < "end")
         )
   );
end;

create trigger check_status_overlap_update
before update on permission_group_has_user
for each row
-- check if the permission_group if of type status
when exists (
   select 1 from permission_groups
   where id = new.permission_group and meta = 'status'
)
begin
   select raise(abort, 'overlapping time span for status assignment is not allowed')
   where exists (
      select 1 
      from permission_group_has_user
      where permission_group = new.permission_group
         and user = new.user
         and (
            (start < new.start and new.start < "end")
            or (start < new."end" and new."end" < "end")
         )
   );
end;