BEGIN ;

CREATE DATABASE IF NOT EXISTS voting;

USE voting;

CREATE TABLE IF NOT EXISTS Users(
    user_id SERIAL NOT NULL ,
    username VARCHAR(150) NOT NULL UNIQUE,
    hashed_password VARCHAR(300) NOT NULL ,
    email VARCHAR(254) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    party_id INTEGER ,
    session TEXT ,
    secret_key TEXT ,
    PRIMARY KEY(user_id),
    CONSTRAINT unique_username_user UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS User_Permissions(
    permission_id INTEGER NOT NULL ,
    user_id INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS Permissions(
    permission_id SERIAL NOT NULL ,
    permission VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Party(
    party_id SERIAL NOT NULL ,
    party VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Candidate(
    candidate_id SERIAL NOT NULL ,
    user_id INTEGER NOT NULL ,
    party_id INTEGER NOT NULL ,
    votes INTEGER DEFAULT 0
);

COMMIT ;

BEGIN ;

INSERT INTO Users(username, hashed_password, email, first_name, last_name, party_id)
VALUES  ('al', 'password', 'axs4986@gmail.com', 'Alberto', 'Serrano', 1),
        ('jam', 'password', 'jxa7578@rit.edu', 'Jahongir', 'Amirkulov', 1),
        ('dan', 'password', 'dxm9604@rit.edu', 'Daniel', 'Moore', 1);

INSERT INTO Party(party)
VALUES  ('democrat'), ('republican'), ('reform'), ('libertarian'), ('socialist'), ('natural'), ('constitution'),
        ('green');

COMMIT ;
