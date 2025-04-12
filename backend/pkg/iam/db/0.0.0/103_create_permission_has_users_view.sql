create view permission_has_users as 
select uf.*, pghp.permission
from user_has_permission_groups uhpg
join permission_group_has_permission pghp on uhpg.permission_group = pghp.permission_group
join user_facades uf on uhpg.user = uf.id;