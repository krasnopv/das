CREATE TABLE IF NOT EXISTS DAS.COMPETITION_PRODUCT_STATUS (
  ID SERIAL NOT NULL  PRIMARY KEY ,
  PRODUCT_STATUS_NAME VARCHAR (16) NOT NULL UNIQUE,
  PRODUCT_STATUS_DESC TEXT,
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW()
);