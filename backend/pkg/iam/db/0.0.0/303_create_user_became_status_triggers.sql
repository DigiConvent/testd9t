create trigger after_insert_user_became_status
after insert on user_became_status
for each row
begin
   insert into permission_group_has_user (permission_group, user, start) values 
      (new.status, new.user, new.start);

   update permission_group_has_user
      set "end" = (
         select "end" 
         from user_status_interval usi
         where permission_group_has_user.permission_group = usi.status and usi.user = permission_group_has_user.user
      )
   where user = new.user and permission_group = new.status;
end;

create trigger after_update_user_became_status 
after update on user_became_status
for each row
begin 
   update permission_group_has_user
      set "end" = (
         select "end" 
         from user_status_interval usi
         where permission_group_has_user.permission_group = usi.status and usi.user = permission_group_has_user.user
      )
   where user = new.user and permission_group = new.status;
end;

create trigger after_delete_user_became_status
after delete on user_became_status
for each row
begin
   delete from permission_group_has_user
   where user = old.user and permission_group = old.status;

   update permission_group_has_user
      set "end" = (
         select "end" 
         from user_status_interval usi
         where permission_group_has_user.permission_group = usi.status and usi.user = permission_group_has_user.user
      )
   where user = old.user and permission_group = old.status;
end;