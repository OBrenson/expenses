CREATE USER myusername WITH PASSWORD 'mypassword';

CREATE SCHEMA IF NOT EXISTS expenses AUTHORIZATION myusername;

ALTER DEFAULT PRIVILEGES IN SCHEMA expenses GRANT ALL PRIVILEGES ON TABLES TO myusername;
ALTER DEFAULT PRIVILEGES IN SCHEMA expenses GRANT USAGE ON SEQUENCES TO myusername;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA expenses TO myusername;

create table if not exists expenses.user (
    id serial primary key,
    username varchar (50) unique,
    is_admin boolean
);

create table if not exists expenses.data (
    user_id integer references expenses.user not null,
    sum float8,
    expense_type varchar(50),
    expense_date timestamp
);

create index if not exists data_index on expenses.data (user_id);