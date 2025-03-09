insert into users (id, emailaddress, first_name, last_name, enabled) values
('00000000-0000-0000-0000-000000000000', '', 'Admin', 'McAdmin', true);

insert into permissions (name) values ('super');
insert into permission_groups (id, name, abbr, is_group, is_node, description, "generated") values ('00000000-0000-0000-0000-000000000000', 'Super', 'super', true, true, 'Super user', true);
insert into permission_group_has_permission (permission_group, permission) values ('00000000-0000-0000-0000-000000000000', 'super');
insert into permission_group_has_user (permission_group, user) values ('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000');