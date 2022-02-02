-- Creating table to store cache data
CREATE TABLE cache(
    key varchar (255) PRIMARY KEY,
    value varchar (255),
    timeout timestamp
);