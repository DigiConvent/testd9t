create table if not exists user_became_status (
  "user" uuid not null references users(id) on delete cascade,
  "status" uuid not null references user_status(id) on delete cascade,
  "date" timestamp not null,
  active boolean default false,
  description varchar default '',
  primary key ("user", "status")
);