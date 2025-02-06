create view user_has_permissions as 
select uhpg.user, pghp.permission
from user_has_permission_groups uhpg
join permission_group_has_permission pghp on uhpg.permission_group = pghp.permission_group;