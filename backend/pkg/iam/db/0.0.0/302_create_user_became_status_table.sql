create table user_became_status (
   user uuid not null references users(id) on delete cascade,
   status uuid not null references user_status(id) on delete cascade,
   start timestamp not null,
   "end" timestamp default null,
   description varchar default '',
   primary key (user, status)
);