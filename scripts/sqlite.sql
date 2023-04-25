drop table if exists users;
create table users(
    id integer primary key autoincrement,
    name varchar(225) not null,
    password varchar(225) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp default current_timestamp
);
insert into users(name, password, created_at, updated_at, deleted_at) 
    values(
        "admin", 
        "$2a$14$hChjXBy1GdHJE5XjvQX5IuCqVfkvowJnT4A4XtNCW1wl8Hp8KMhvm",
        datetime(),
        null,
        null
    );