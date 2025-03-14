create table products (
    id serial primary key,
    user_id integer,
    title varchar not null,
    description text not null,
    price integer not null default 0,
    is_available boolean not null default false
);

create table users (
    id serial primary key,
    name varchar not null,
    email varchar not null,
    password_hash varchar not null
);