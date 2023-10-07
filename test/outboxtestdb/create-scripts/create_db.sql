CREATE DATABASE OutboxTestDb;

USE OutboxTestDb;

CREATE USER 'AdminUser'@'%' IDENTIFIED BY 'AdminPassword';
GRANT ALL PRIVILEGES ON *.* TO 'AdminUser'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;

SHOW GRANTS FOR 'AdminUser'@'%';

CREATE TABLE IF NOT EXISTS any_tables(
    `id` varchar(36) NOT NULL,
    `any_data` varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARACTER SET = UTF8MB4
    COLLATE = utf8mb4_unicode_520_ci;

CREATE TABLE IF NOT EXISTS outboxes(
    `id` varchar(36) NOT NULL,
    `table_id` varchar(36) NOT NULL,
    `table_name` varchar(100) NOT NULL,
    `status` varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARACTER SET = UTF8MB4
    COLLATE = utf8mb4_unicode_520_ci;

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;

INSERT INTO any_tables(id, any_data) VALUES ('c58b4e85-600b-4171-91ff-80af251ddea0', 'data1');
INSERT INTO any_tables(id, any_data) VALUES ('c58b4e85-600b-4171-91ff-80af251ddea1', 'data2');
INSERT INTO any_tables(id, any_data) VALUES ('c58b4e85-600b-4171-91ff-80af251ddea2', 'data3');

INSERT INTO outboxes(id, table_id, table_name, status) VALUES ('bbbbbbbb-600b-4171-91ff-80af251ddea0', 'c58b4e85-600b-4171-91ff-80af251ddea0', 'any_tables', 'created');
INSERT INTO outboxes(id, table_id, table_name, status) VALUES ('bbbbbbbb-600b-4171-91ff-80af251ddea1', 'c58b4e85-600b-4171-91ff-80af251ddea1', 'any_tables', 'created');
INSERT INTO outboxes(id, table_id, table_name, status) VALUES ('bbbbbbbb-600b-4171-91ff-80af251ddea2', 'c58b4e85-600b-4171-91ff-80af251ddea2', 'any_tables', 'created');

SET GLOBAL binlog_format='ROW';
SET SESSION binlog_format='ROW';