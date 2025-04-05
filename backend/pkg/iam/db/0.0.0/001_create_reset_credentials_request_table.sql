create table reset_credentials_requests (
   user uuid not null,
   token text not null,
   created_at timestamp default current_timestamp,
   primary key (user)
);