DROP DATABASE IF EXISTS StockManagment;
CREATE DATABASE StockManagment;

USE StockManagment;

CREATE TABLE Users (
    id varchar(200),
    first_name varchar(200),
    last_name varchar(200),
    email varchar(200),
    password varchar(200),
    img varchar(200),
    creation_date datetime,

    PRIMARY KEY (id)
);