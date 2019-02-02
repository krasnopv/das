CREATE TABLE IF NOT EXISTS DAS.EVENT_ROUND_SCHEDULE (
  ID SERIAL NOT NULL PRIMARY KEY,
  EVENT_ID INTEGER REFERENCES DAS.EVENT(ID),
  ROUND_ID INTEGER NOT NULL,
  PRELIMINARY_ROUND_IND BOOLEAN DEFAULT TRUE,
  ESTIMATED_STARTTIME TIMESTAMP DEFAULT NOW(),
  ACTUAL_STARTTIME TIMESTAMP DEFAULT NOW(),
  CREATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT(ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER REFERENCES DAS.ACCOUNT(ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (EVENT_ID, ROUND_ID)
);
CREATE INDEX ON DAS.EVENT_ROUND_SCHEDULE (EVENT_ID);
