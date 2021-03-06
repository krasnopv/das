-- create table das.federation
CREATE TABLE IF NOT EXISTS DAS.FEDERATION (
  ID SERIAL NOT NULL PRIMARY KEY ,
  NAME VARCHAR (64) NOT NULL UNIQUE,
  ABBREVIATION VARCHAR (10) UNIQUE,
  DESCRIPTION TEXT,
  YEAR_FOUNDED INTEGER NOT NULL,
  COUNTRY_ID INTEGER NOT NULL REFERENCES DAS.COUNTRY(ID),
  CREATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT (ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT (ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW()
);

INSERT INTO DAS.FEDERATION (NAME, ABBREVIATION, YEAR_FOUNDED, COUNTRY_ID) VALUES ('Canada DanceSport','CDS',1900, (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.FEDERATION (NAME, ABBREVIATION, YEAR_FOUNDED, COUNTRY_ID) VALUES ('Collegiate','COL',1900, (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.FEDERATION (NAME, ABBREVIATION, YEAR_FOUNDED, COUNTRY_ID) VALUES ('Independent/Unaffiliated','IND',1900, (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.FEDERATION (NAME, ABBREVIATION, YEAR_FOUNDED, COUNTRY_ID) VALUES ('National Dance Council of America Inc.','NDCA',1948, (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.FEDERATION (NAME, ABBREVIATION, YEAR_FOUNDED, COUNTRY_ID) VALUES ('National Dance Council of Canada','NDCC',1900, (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'CAN'));
INSERT INTO DAS.FEDERATION (NAME, ABBREVIATION, YEAR_FOUNDED, COUNTRY_ID) VALUES ('USA Dance Inc.','USA DANCE',1965, (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'USA'));
INSERT INTO DAS.FEDERATION (NAME, ABBREVIATION, YEAR_FOUNDED, COUNTRY_ID) VALUES ('World DanceSport Federation','WDSF',1957, (SELECT C.ID FROM DAS.COUNTRY C WHERE C.ABBREVIATION = 'SUI'));