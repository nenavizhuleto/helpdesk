CREATE TABLE IF NOT EXISTS companies (
    id TEXT NOT NULL, -- Company INN
    name TEXT NOT NULL,
    slug TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS branches (
    id TEXT NOT NULL, -- UUID
    company_id TEXT NOT NULL,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    contacts TEXT NOT NULL,
    description TEXT NOT NULL,
    PRIMARY KEY(id)
);


CREATE TABLE IF NOT EXISTS users (
    id TEXT NOT NULL, -- UUID
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS devices (
    ip TEXT NOT NULL,
    user_id TEXT,
    network_id TEXT,
    type TEXT DEFAULT "Unknown" NOT NULL, -- PC, Terminal, Router
    PRIMARY KEY(ip)
);

CREATE TABLE IF NOT EXISTS networks (
    netmask TEXT NOT NULL,
    branch_id TEXT NOT NULL,
    PRIMARY KEY(netmask)
);

CREATE TABLE IF NOT EXISTS tasks (
    id TEXT NOT NULL,
    human_number INTEGER NOT NULL,
    user_id TEXT NOT NULL,
    name TEXT NOT NULL,
    subject TEXT NOT NULL,
    status TEXT NOT NULL,
    time_created TEXT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS comments (
    id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    subject TEXT NOT NULL,
    content TEXT NOT NULL,
    time_created TEXT NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO companies VALUES ("0000000000", "Development Environment Company", "devenvcomp");
INSERT INTO branches VALUES ("0000-1111-0000-1111", "0000000000", "Dev Branch", "Development Space", "dev@env.com", "Master Branch");
INSERT INTO networks VALUES ("127.0.0.1/32", "0000-1111-0000-1111");

INSERT INTO companies VALUES ("gtINN", "General Telecom", "gt");
INSERT INTO branches VALUES ("gtUUID", "gtINN", "Основной отдел", "Труда 24", "880005553535", "Главный офис");
INSERT INTO networks VALUES ("172.16.222.0/24", "gtUUID");


INSERT INTO companies VALUES ("7451195700", "Южноуральский Лизинговый Центр", "lizing");
INSERT INTO branches VALUES ("0", "7451195700", "Челябинск", "неизвестно", "неизвестно", "Челябинск");
INSERT INTO networks VALUES ("172.28.132.0/25", "0");
-- Уфа 172.28.132.128/27
-- 172.28.132.129
-- 81.30.218.196
-- Wkb4Sgz3MUNv7NNT
INSERT INTO branches VALUES ("1", "7451195700", "Уфа", "неизвестно", "неизвестно", "Уфа");
INSERT INTO networks VALUES ("172.28.132.128/27", "1");
-- Тюмень 172.28.132.160/27
-- 172.28.132.161
-- 81.30.218.196
-- jSP86bb5sqHvtfjs
INSERT INTO branches VALUES ("2", "7451195700", "Тюмень", "неизвестно", "неизвестно", "Тюмень");
INSERT INTO networks VALUES ("172.28.132.160/27", "2");
-- Казань 172.28.132.192/27                             
-- 172.28.132.193                                     
-- 91.245.36.246
-- ng7nTknQUGhOn8QJ
INSERT INTO branches VALUES ("3", "7451195700", "Казань", "неизвестно", "неизвестно", "Казань");
INSERT INTO networks VALUES ("172.28.132.192/27", "3");
-- Екат 172.28.132.224/27
-- 172.28.132.225
-- 91.194.244.26
-- bwI3sBevXWv6sUHa
INSERT INTO branches VALUES ("4", "7451195700", "Екат", "неизвестно", "неизвестно", "Екат");
INSERT INTO networks VALUES ("172.28.132.224/27", "4");
-- Самара 172.28.133.0/27
-- 172.28.133.1
-- 193.27.241.8
-- 2xAwlJBLkpg2Am4o
INSERT INTO branches VALUES ("5", "7451195700", "Самара", "неизвестно", "неизвестно", "Самара");
INSERT INTO networks VALUES ("172.28.133.0/27", "5");
-- Новосибирск 172.28.133.32/27
-- 172.28.133.33
-- 109.202.20.216
-- 57KkjSewNtfCGsyY
INSERT INTO branches VALUES ("6", "7451195700", "Новосибирск", "неизвестно", "неизвестно", "Новосибирск");
INSERT INTO networks VALUES ("172.28.133.32/27", "6");
-- Краснодар 72.28.133.64/27
-- 172.28.133.65
-- 185.105.171.226
-- vpQBVx0Sew9SCJCl
INSERT INTO branches VALUES ("7", "7451195700", "Краснодар", "неизвестно", "неизвестно", "Краснодар");
INSERT INTO networks VALUES ("172.28.133.64/27", "7");
-- СПБ 172.28.133.128/2
-- 172.28.133.129
-- 193.105.107.143
-- NTGLBBBkrJcgUjgs
INSERT INTO branches VALUES ("8", "7451195700", "СПБ", "неизвестно", "неизвестно", "СПБ");
INSERT INTO networks VALUES ("172.28.133.128/27", "8");
-- МСК 172.28.133.96/27
-- 172.28.133.97
-- 212.193.99.28
-- 1YIOmMd2lJntEWDV
INSERT INTO branches VALUES ("9", "7451195700", "МСК", "неизвестно", "неизвестно", "МСК");
INSERT INTO networks VALUES ("172.28.133.96/27", "9");
-- Н.Новгород
-- 172.28.133.160/27
-- 93.120.238.18
-- JmwHMUlGHsXw6MeT
INSERT INTO branches VALUES ("10", "7451195700", "Нижний Новгород", "неизвестно", "неизвестно", "Нижний Новгород");
INSERT INTO networks VALUES ("172.28.133.160/27", "10");
