create trigger after_insert_user_became_role
after insert on user_became_role
for each row
begin
   insert into permission_group_has_user (permission_group, user, start) values 
      (new.role, new.user, new.start);

   update permission_group_has_user
      set "end" = (
         select "end" 
         from user_role_interval uri
         where permission_group_has_user.permission_group = uri.role and uri.user = permission_group_has_user.user
      )
   where user = new.user and permission_group = new.role;
end;

create trigger after_update_user_became_role 
after update on user_became_role
for each row
begin 
   update permission_group_has_user
      set "end" = (
         select "end" 
         from user_role_interval uri
         where permission_group_has_user.permission_group = uri.role and uri.user = permission_group_has_user.user
      )
   where user = new.user and permission_group = new.role;
end;

create trigger after_delete_user_became_role
after delete on user_became_role
for each row
begin
   delete from permission_group_has_user
   where user = old.user and permission_group = old.role;

   update permission_group_has_user
      set "end" = (
         select "end" 
         from user_role_interval uri
         where permission_group_has_user.permission_group = uri.role and uri.user = permission_group_has_user.user
      )
   where user = old.user and permission_group = old.role;
end;