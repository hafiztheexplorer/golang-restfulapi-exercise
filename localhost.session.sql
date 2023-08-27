/* create the database */
USE gorestfulapi_exercise;

CREATE TABLE category(
    id int PRIMARY KEY AUTO_INCREMENT,
    namakategori VARCHAR(255)NOT NULL
)ENGINE=InnoDB;

SELECT * FROM category;

