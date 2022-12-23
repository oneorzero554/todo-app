CREATE TABLE list
(
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255) null
);

CREATE TABLE item
(
    id serial not null unique,
    name varchar(255) not null,
    list_id int references list (id) on delete cascade not null,
    description varchar(255) null,
    is_done boolean not null default false
);