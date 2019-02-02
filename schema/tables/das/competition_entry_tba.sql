CREATE TABLE IF NOT EXISTS DAS.COMPETITION_ENTRY_TBA (
  ID SERIAL NOT NULL PRIMARY KEY,
  ACCOUNT_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  COMPETITION_ID INTEGER NOT NULL REFERENCES DAS.COMPETITION (ID),
  ROLE TEXT NOT NULL,
  PROFICIENCY TEXT NOT NULL,
  STYLE TEXT NOT NULL,
  HEIGHT TEXT NOT NULL,
  SCHOOL TEXT NOT NULL,
  CONTACT_EMAIL TEXT,
  CONTACT_PHONE TEXT,
  MISC_INFO TEXT,
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT(ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT(ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX ON DAS.COMPETITION_ENTRY_TBA (ACCOUNT_ID);
CREATE INDEX ON DAS.COMPETITION_ENTRY_TBA (COMPETITION_ID);