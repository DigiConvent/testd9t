create table configurations (
    domain text default '',
    telegram_bot_token text default ''
);

insert into configurations (domain, telegram_bot_token) values ('', '');