create trigger after_insert_user_status
after insert on user_status
for each row
begin
   insert into permission_groups (id, name, abbr, is_group, is_node, meta, description, generated) values 
      (new.id, new.name, new.abbr, 1, 1, 'status', new.description, 1);
end;

create trigger after_update_user_status
after update on user_status
for each row
begin
   update permission_groups set name = new.name, abbr = new.abbr, description = new.description where id = old.id;
end;

create trigger after_delete_user_status
after delete on user_status
for each row
begin
   delete from permission_groups where id = old.id;
end;
