create database blog;

create table blog.t_article(
    id int auto_increment primary key,
    title varchar(250),
    content text ,
    `tag` varchar(200),
    createdTime datetime not null default now()
)engine=innodb charset=utf8mb4;