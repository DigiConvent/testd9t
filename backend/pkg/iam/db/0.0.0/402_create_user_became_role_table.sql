create table user_became_role (
   user uuid not null references users(id) on delete cascade,
   role uuid not null references user_roles(id) on delete cascade,
   start timestamp not null,
   "end" timestamp not null,
   description varchar default ''
);