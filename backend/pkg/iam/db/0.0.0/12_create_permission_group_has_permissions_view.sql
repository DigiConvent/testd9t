create view permission_group_has_permissions as
select p.id, p.name, pghpgs.id as permission_group_id, pghpgs.name as permission_group_name
from permission_group_has_permission_groups pghpgs
left join permission_group_has_permission pghp on pghpgs.id = pghp.permission_group;