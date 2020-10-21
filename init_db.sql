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
    party_id INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS Votes(
    candidate_id INTEGER NOT NULL ,
    votes INTEGER DEFAULT 0
);


COMMIT ;
