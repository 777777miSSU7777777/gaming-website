CREATE DATABASE GAMING_WEBSITE;

USE GAMING_WEBSITE;

CREATE TABLE USERS (
    USER_ID BIGINT NOT NULL AUTO_INCREMENT,
    USERNAME NVARCHAR(50) NOT NULL,
    BALANCE BIGINT NOT NULL,
    PRIMARY KEY(USER_ID),
    CHECK(USERNAME <> ""),
    CHECK(BALANCE >= 0)
);