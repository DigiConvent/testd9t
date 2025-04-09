create table configurations (
    domain text default '',
    telegram_bot_token text default '',
    maintenance boolean default false
);

insert into configurations (domain, telegram_bot_token) values ('', '');