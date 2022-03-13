create table articles
(
    id         int unsigned auto_increment
        primary key,
    created_at datetime     not null,
    updated_at datetime     not null,
    deleted_at datetime     null,
    title      varchar(255) not null,
    body       text         not null
);

INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (1, '2022-03-11 20:56:31', '2022-03-11 20:56:31', '2022-03-11 22:28:52', 'My article #0', 'Article body #0');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (2, '2022-03-11 20:56:31', '2022-03-11 20:56:31', null, 'My article #1', 'Article body #1');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (3, '2022-03-11 20:56:31', '2022-03-11 20:56:31', null, 'My article #2', 'Article body #2');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (4, '2022-03-11 20:56:31', '2022-03-11 20:56:31', null, 'My article #3', 'Article body #3');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (5, '2022-03-11 20:56:31', '2022-03-11 20:56:31', null, 'My article #4', 'Article body #4');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (6, '2022-03-11 20:56:31', '2022-03-11 20:56:31', null, 'My article #5', 'Article body #5');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (7, '2022-03-11 20:56:31', '2022-03-11 20:56:31', null, 'My article #6', 'Article body #6');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (8, '2022-03-11 20:56:31', '2022-03-11 20:56:31', null, 'My article #7', 'Article body #7');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (9, '2022-03-11 20:56:31', '2022-03-11 20:56:31', null, 'My article #8', 'Article body #8');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (10, '2022-03-11 20:56:31', '2022-03-11 20:56:31', null, 'My article #9', 'Article body #9');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (11, '2022-03-12 10:40:44', '2022-03-13 11:57:43', null, 'Title New #?+5', 'Some body');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (12, '2022-03-12 10:48:34', '2022-03-12 12:01:13', '2022-03-13 00:02:55', 'Title New #2+1+1', 'Some body + 1');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (13, '2022-03-13 00:57:32', '2022-03-13 00:57:32', null, 'New from xml #1', 'New super body');
INSERT INTO articles (id, created_at, updated_at, deleted_at, title, body) VALUES (14, '2022-03-13 00:57:51', '2022-03-13 00:57:51', '2022-03-13 11:51:43', 'New from xml #2', 'New super body');
