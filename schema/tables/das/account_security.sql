CREATE TABLE IF NOT EXISTS DAS.ACCOUNT_SECURITY (
  ID SERIAL NOT NULL PRIMARY KEY,
  ACCOUNT_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT(ID)
);