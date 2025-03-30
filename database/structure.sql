-- Create user if it doesn't exist
DO
$$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'goapi') THEN
        CREATE ROLE goapi WITH LOGIN PASSWORD 'password';
    END IF;
END
$$;

-- Create database if it doesn't exist
DO
$$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'goapi') THEN
        CREATE DATABASE goapi OWNER goapi;
    END IF;
END
$$;

\c goapi;

CREATE TABLE IF NOT EXISTS products (
    id serial primary key,
    user_id integer,
    title varchar not null,
    description text not null,
    price integer not null default 0,
    is_available boolean not null default false
);

CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    name varchar not null,
    email varchar not null UNIQUE,
    password_hash varchar not null
);