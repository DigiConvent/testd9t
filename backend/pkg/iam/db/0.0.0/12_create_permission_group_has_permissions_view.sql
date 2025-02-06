create view permission_group_has_permissions as
select pghp.permission as name, pghpgs.id as permission_group, pghpgs.name as permission_group_name
from permission_group_has_permission_groups pghpgs
join permission_group_has_permission pghp on pghpgs.id = pghp.permission_group;