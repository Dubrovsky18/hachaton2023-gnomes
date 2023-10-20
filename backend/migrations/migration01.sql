CREATE DATABASE test ENCODING 'UTF8';

CREATE TABLE teacher (
    id              SERIAL PRIMARY KEY,
    first_name      VARCHAR(50) NOT NULL,
    second_name     VARCHAR(50) NOT NULL,
    phone           VARCHAR(20) NOT NULL,
    email           VARCHAR(20) NOT NULL,
    birthday        DATE NOT NULL,
    pass            VARCHAR(255) NOT NULL,
    week_hours      INTEGER NOT NULL,
    weekends_day    VARCHAR(25) NOT NULL,
);

CREATE TABLE admin (
    id              SERIAL PRIMARY KEY,
    first_name      VARCHAR(50) NOT NULL,
    second_name     VARCHAR(50) NOT NULL,
    phone           VARCHAR(20) NOT NULL,
    email           VARCHAR(20) NOT NULL,
    birthday        DATE NOT NULL,
    pass            VARCHAR(255) NOT NULL,
);

