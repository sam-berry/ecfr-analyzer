CREATE TABLE agency
(
    id               SERIAL PRIMARY KEY,
    agencyId         UUID UNIQUE NOT NULL,
    name             TEXT        NOT NULL,
    shortName        TEXT        NOT NULL,
    sortableName     TEXT        NOT NULL,
    slug             TEXT UNIQUE NOT NULL,
    createdTimestamp TIMESTAMP   NOT NULL
);
