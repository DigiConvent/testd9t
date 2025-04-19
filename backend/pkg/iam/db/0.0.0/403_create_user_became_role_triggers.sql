create trigger before_insert_user_became_role
before insert on user_became_role
for each row
begin
   select case
      when exists (
         select 1 
         from user_became_role
         where user = new.user
         and role = new.role
         and (
            (new.start between start and "end") or
            (new.end between start and "end")
         )
      )
      then raise(abort, 'user already has this role with overlapping time period')
   end;
end;

create trigger after_insert_user_became_role
after insert on user_became_role
for each row
begin
   insert into permission_group_has_user (permission_group, user, start, "end") values 
      (new.role, new.user, new.start, new.end);
end;

create trigger after_update_user_became_role 
after update on user_became_role
for each row
begin 
   update permission_group_has_user
   set start = new.start, end = new.end
   where user = new.user and permission_group = new.role;
end;

create trigger after_delete_user_became_role
after delete on user_became_role
for each row
begin
   delete from permission_group_has_user
   where user = old.user and permission_group = old.role;
end;