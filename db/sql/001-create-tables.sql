CREATE DATABASE sample_db;
USE sample_db;

---- drop ----
DROP TABLE IF EXISTS `t_task`;

---- create ----
create table IF not exists `t_task`
(
 `id`               INT(11) AUTO_INCREMENT,
 `task`             VARCHAR(256) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

---- insert ----
INSERT INTO t_task VALUES (1,'tanaka', '2019-10-04 15:25:07', '2019-10-04 15:25:07');
INSERT INTO t_task VALUES (2,'kage', '2020-04-02 15:25:07', '2020-04-04 15:25:07');
INSERT INTO t_task VALUES (3,'yamada', '2020-05-05 15:25:07', '2020-05-06 15:25:07');
