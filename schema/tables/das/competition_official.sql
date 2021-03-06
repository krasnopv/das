-- Table for invitation to competitions.
CREATE TABLE IF NOT EXISTS DAS.COMPETITION_OFFICIAL_INVITATION (
  ID SERIAL NOT NULL PRIMARY KEY,
  ORGANIZER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID), -- must be an organizer
  RECIPIENT_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  COMPETITION_ID INTEGER NOT NULL REFERENCES DAS.COMPETITION (ID),
  ROLE_ASSIGNED INTEGER NOT NULL REFERENCES DAS.ACCOUNT_ROLE (ID), -- recipient must have a role
  MESSAGE TEXT,
  STATUS VARCHAR (16) NOT NULL, -- values are: Accepted, Rejected, and Pending
  EXPIRATION_DATE DATE NOT NULL DEFAULT NOW(),
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT(ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT(ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  CHECK (ORGANIZER_ID != RECIPIENT_ID),
  UNIQUE (ORGANIZER_ID, RECIPIENT_ID, COMPETITION_ID, ROLE_ASSIGNED, STATUS)
);

CREATE INDEX ON DAS.COMPETITION_OFFICIAL_INVITATION (ORGANIZER_ID);
CREATE INDEX ON DAS.COMPETITION_OFFICIAL_INVITATION (COMPETITION_ID);
CREATE INDEX ON DAS.COMPETITION_OFFICIAL_INVITATION (CREATE_USER_ID);
CREATE INDEX ON DAS.COMPETITION_OFFICIAL_INVITATION (UPDATE_USER_ID);

CREATE TABLE IF NOT EXISTS DAS.COMPETITION_OFFICIAL (
  ID SERIAL NOT NULL PRIMARY KEY,
  COMPETITION_ID INTEGER NOT NULL REFERENCES DAS.COMPETITION (ID),
  OFFICIAL_ACCOUNT_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  OFFICIAL_ROLE_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT_ROLE (ID),
  EFFECTIVE_FROM DATE NOT NULL DEFAULT NOW(),
  EFFECTIVE_UNTIL DATE NOT NULL DEFAULT NOW(),
  ASSIGNED_BY INTEGER NOT NULL REFERENCES DAS.ACCOUNT (ID),
  CREATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT(ID),
  DATETIME_CREATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UPDATE_USER_ID INTEGER NOT NULL REFERENCES DAS.ACCOUNT(ID),
  DATETIME_UPDATED TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (COMPETITION_ID, OFFICIAL_ACCOUNT_ID, OFFICIAL_ROLE_ID),
  CHECK ( EFFECTIVE_UNTIL > EFFECTIVE_FROM ),
  CHECK ( DATETIME_CREATED <= DATETIME_UPDATED)
);

CREATE INDEX ON DAS.COMPETITION_OFFICIAL (COMPETITION_ID);
CREATE INDEX ON DAS.COMPETITION_OFFICIAL (OFFICIAL_ACCOUNT_ID);
CREATE INDEX ON DAS.COMPETITION_OFFICIAL (OFFICIAL_ROLE_ID);
CREATE INDEX ON DAS.COMPETITION_OFFICIAL (CREATE_USER_ID);
CREATE INDEX ON DAS.COMPETITION_OFFICIAL (UPDATE_USER_ID);