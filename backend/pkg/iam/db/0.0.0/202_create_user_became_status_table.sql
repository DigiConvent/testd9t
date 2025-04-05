create table user_became_status (
   "user" uuid not null references users(id) on delete cascade,
   "status" uuid not null references user_status(id) on delete cascade,
   "start" timestamp not null,
   description varchar default '',
   primary key ("user", "status")
);

create view user_status_interval as
select 
   "user",
   status,
   start,
   coalesce(
      datetime(lead(start) over (partition by user order by start), '-1 second'),
      null
   ) as end
from user_became_status;