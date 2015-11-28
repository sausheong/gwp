drop table posts;

create table posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);