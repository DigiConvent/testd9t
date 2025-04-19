create trigger after_insert_user_role
after insert on user_roles
for each row
begin
   insert into permission_groups (id, name, abbr, meta, description, generated) values 
      (new.id, new.name, new.abbr, 'role', new.description, 1);
end;

create trigger after_update_user_role
after update on user_roles
for each row
begin
   update permission_groups set name = new.name, abbr = new.abbr, description = new.description where id = old.id;
end;

create trigger after_delete_user_role
after delete on user_roles
for each row
begin
   delete from permission_groups where id = old.id;
end;
