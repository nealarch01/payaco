-- WARNING: running this script will cause your previous database to drop
DROP DATABASE IF EXISTS development;
CREATE DATABASE development;

-- Use the development database
\c development;

-- Table is named 'account' because 'user' is a reserved word
CREATE TABLE account (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(32) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(65) NOT NULL,
    BALANCE DECIMAL(12, 2) NOT NULL
);


insert into account (id, username, first_name, last_name, phone_number, email, password, balance) values (1, 'lharrington0', 'Lucky', 'Harrington', '+48-457-961-4950', 'lharrington0@ameblo.jp', 'fkVrJvFI', 168.58);
insert into account (id, username, first_name, last_name, phone_number, email, password, balance) values (2, 'hoxnam1', 'Helenelizabeth', 'Oxnam', '+7-720-764-4230', 'hoxnam1@ovh.net', 'lzrAXBG', 283.57);
insert into account (id, username, first_name, last_name, phone_number, email, password, balance) values (3, 'glemmers2', 'Gilberto', 'Lemmers', '+62-336-330-8967', 'glemmers2@issuu.com', 'NaAZeF', 208.88);
insert into account (id, username, first_name, last_name, phone_number, email, password, balance) values (4, 'eboffin3', 'Ellynn', 'Boffin', '+62-697-566-2893', 'eboffin3@latimes.com', 'IhjbH4R', 206.95);
insert into account (id, username, first_name, last_name, phone_number, email, password, balance) values (5, 'kbidgood4', 'Karlene', 'Bidgood', '+84-860-596-6871', 'kbidgood4@addthis.com', 'zjwFmrbT', 26.9);
insert into account (id, username, first_name, last_name, phone_number, email, password, balance) values (6, 'pschroter5', 'Patricio', 'Schroter', '+86-213-785-8233', 'pschroter5@imdb.com', 'IeaxWT', 121.72);
insert into account (id, username, first_name, last_name, phone_number, email, password, balance) values (7, 'idunmuir6', 'Ida', 'Dunmuir', '+60-824-473-1470', 'idunmuir6@mlb.com', '5jYPZ0um9fef', 252.07);
insert into account (id, username, first_name, last_name, phone_number, email, password, balance) values (8, 'sofeeney7', 'Siward', 'O''Feeney', '+375-711-917-6656', 'sofeeney7@google.co.jp', 'wncLGW', 220.69);
insert into account (id, username, first_name, last_name, phone_number, email, password, balance) values (9, 'rmartyntsev8', 'Ralina', 'Martyntsev', '+33-839-643-9990', 'rmartyntsev8@ow.ly', 'EzJiIRVRVc', 254.98);
insert into account (id, username, first_name, last_name, phone_number, email, password, balance) values (10, 'dmannock9', 'Drucy', 'Mannock', '+63-746-396-7160', 'dmannock9@deliciousdays.com', 'P2BXRx', 195.88);


CREATE TABLE transactions (
    sender INT REFERENCES account(id),
    receiver INT REFERENCES account(id),
    amount DECIMAL(12, 2)
);


-- Table that blacklists jwt tokens
CREATE TABLE blacklist (
    token VARCHAR(500) NOT NULL PRIMARY KEY
);
