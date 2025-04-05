create table user_became_role (
   "user" uuid not null references users(id) on delete cascade,
   "role" uuid not null references user_roles(id) on delete cascade,
   "start" timestamp not null,
   description varchar default '',
   primary key ("user", "role")
);

create view user_role_interval as
select 
   "user",
   "role",
   start,
   coalesce(
      datetime(lead(start) over (partition by user order by start), '-1 second'),
      null
   ) as end
from user_became_role;