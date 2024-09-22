DROP DATABASE IF EXISTS StockManagment;
CREATE DATABASE StockManagment;

USE StockManagment;

CREATE TABLE Companies (
    id varchar(200),
    name varchar(200),

    PRIMARY KEY (id)
);

CREATE TABLE Users (
    id varchar(200),
    first_name varchar(200),
    last_name varchar(200),
    email varchar(200),
    password varchar(200),
    img varchar(200),
    creation_date datetime,
    company_id varchar(200),

    PRIMARY KEY (id),
    FOREIGN KEY (company_id) REFERENCES Companies(id)
);

CREATE TABLE Warehouses (
    id varchar(200),
    name varchar(200),
    creation_date datetime,
    creator_id varchar(200),
    company_id varchar(200),

    PRIMARY KEY (id),
    FOREIGN KEY (company_id) REFERENCES Companies(id)
);

CREATE TABLE Colors (
    id varchar(200),
    name varchar(200),
    creation_date datetime,
    creator_id varchar(200),
    color varchar(10),
    company_id varchar(200),
    
    PRIMARY KEY (id),
    FOREIGN KEY (company_id) REFERENCES Companies(id)
);

CREATE TABLE Unitys (
    id varchar(200),
    name varchar(200),
    description varchar(10),
    creation_date datetime,
    creator_id varchar(200),
    color varchar(10),
    company_id varchar(200),
    
    PRIMARY KEY (id),
    FOREIGN KEY (company_id) REFERENCES Companies(id)
);

CREATE TABLE Stuf (
    id varchar(200),
    name varchar(200),
    creation_date datetime,
    creator_id varchar(200),
    company_id varchar(200),
    color_id varchar(200),
    unity_id varchar(200),

    PRIMARY KEY (id),
    FOREIGN KEY (company_id) REFERENCES Companies(id),
    FOREIGN KEY (color_id) REFERENCES Colors(id),
    FOREIGN KEY (unity_id) REFERENCES Unitys(id)
);

CREATE TABLE Stuff_Batch (
    id varchar(200),
    name varchar(200),
    creation_date datetime,
    creator_id varchar(200),
    stuf_id varchar(200),

    PRIMARY KEY (id),
    FOREIGN KEY (stuf_id) REFERENCES Stuf(id)
);

CREATE TABLE Warehouses_Stuff (
    id varchar(200),
    wear_house_id varchar(200),
    stuf_id varchar(200),
    quantity int,
    batch_id varchar(200),

    PRIMARY KEY (id),
    FOREIGN KEY (stuf_id) REFERENCES Stuf(id),
    FOREIGN KEY (batch_id) REFERENCES Stuff_Batch(id),
    FOREIGN KEY (wear_house_id) REFERENCES Warehouses(id)
);