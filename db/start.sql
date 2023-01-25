CREATE SCHEMA IF NOT EXISTS expenses;

create table if not exists expenses.user (
    id serial primary key,
    username varchar (50) unique,
    is_admin boolean
);

create table if not exists expenses.data (
    user_id integer references expenses.user,
    sum money,
    expense_type varchar(50),
    expense_date timestamp
);

create index if not exists data_index on expenses.data (user_id);