CREATE TABLE IF NOT EXISTS COLLEGIATE.COMPETITIVE_EVENT (
  ID SERIAL NOT NULL  PRIMARY KEY ,
  EVENT_ID INTEGER NOT NULL REFERENCES DAS.COMPETITIVE_EVENT(ID),
  SKILL_ID INTEGER NOT NULL REFERENCES DAS.PROFICIENCY(ID),
  STYLE_ID INTEGER NOT NULL REFERENCES DAS.STYLE(ID),
  RANKED BOOLEAN NOT NULL DEFAULT FALSE,
  NEWCOMER_ONLY BOOLEAN NOT NULL DEFAULT FALSE,
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW()
);