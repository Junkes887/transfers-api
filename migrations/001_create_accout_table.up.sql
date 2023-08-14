CREATE TABLE ACCOUNTS (
	ID varchar(100) NOT NULL,
    NAME VARCHAR(200) NOT NULL,
    CPF VARCHAR(200) UNIQUE NOT NULL,
    SECRET VARBINARY(200) NOT NULL,
	BALANCE FLOAT NOT NULL,
    CREATED_AT DATETIME NOT NULL,
    PRIMARY KEY (ID)
);