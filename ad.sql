CREATE DATABASE IF NOT EXISTS dcard;

USE dcard;
CREATE TABLE Ad (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    startAt DATETIME NOT NULL,
    endAt DATETIME NOT NULL,
    ageStart INT NOT NULL,
    ageEnd INT NOT NULL,
    gender INT NOT NULL
);


CREATE TABLE Country (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);


CREATE TABLE Platform (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);



CREATE TABLE Ad_Country (
    adId INT,
    countryId INT,
    FOREIGN KEY (adId) REFERENCES Ad(id),
    FOREIGN KEY (countryId) REFERENCES Country(id)
);


CREATE TABLE Ad_Platform (
    adId INT,
    platformId INT,
    FOREIGN KEY (adId) REFERENCES Ad(id),
    FOREIGN KEY (platformId) REFERENCES Platform(id)
);
