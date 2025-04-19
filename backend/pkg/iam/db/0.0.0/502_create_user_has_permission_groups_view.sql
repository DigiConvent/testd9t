create view user_has_permission_groups as
select 
   distinct pghpga.id as permission_group,
   pghu.user,
   pghpga.implied,
   pghpga.parent
from permission_group_has_user pghu
join permission_group_has_permission_group_ancestors pghpga on pghu.permission_group = pghpga.root;