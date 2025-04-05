create table if not exists user_roles (
   id uuid primary key not null,
   name varchar not null unique,
   abbr varchar not null,
   description varchar default ''
);