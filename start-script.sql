create database if not exists `test_jungle`;
use `test_jungle`;

drop database if exists `users`;

create table `users`
(
    `id`            int primary key auto_increment,
    `username`      varchar(255) not null,
    `password_hash` varchar(255) not null
);

insert into users(username, password_hash)
values ('test', '$2a$10$ZxVBwjZOqBXh52TTyGrS/eCySKcr.cPN0Vk2IAhGWl4LJqz7ZWUI6');

create table `images`
(
    `id`         int primary key auto_increment,
    `user_id`    int not null,
    `image_path` varchar(2000),
    `image_url`  varchar(2000),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);
