-- Active: 1696642727227@@127.0.0.1@5432@sell_product@public

CREATE Table users (
    id BIGINT PRIMARY key,
    name VARCHAR(20),
    age int,
    address VARCHAR(20)
);

SELECT *FROM users;