insert into users (id, emailaddress, first_name, last_name, enabled) values
('00000000-0000-0000-0000-000000000000', '', 'Admin', 'McAdmin', 1);

insert into permissions (name, description) values ('admin', 'Permission to bypass all permissions.'); -- this is the admin user
insert into permission_groups (id, name, abbr, description, meta) values ('00000000-0000-0000-0000-000000000000', 'admin', 'admin', 'A role for bypassing all permissions', 'role');
insert into permission_group_has_user (permission_group, user, start, "end", comment) values ('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', datetime('now', 'localtime'), datetime('9999-12-31T23:59:59'), '');
insert into permission_group_has_permission (permission_group, permission) values ('00000000-0000-0000-0000-000000000000', 'admin');
