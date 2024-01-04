CREATE TABLE if not exists Server (
    "TimePosted" INTEGER NOT NULL,
    "NextPostTime" INTEGER NOT NULL,
    "ServerTime" INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS Players (
    "Id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "Rank" INTEGER NOT NULL,
    "Name" INTEGER NOT NULL,
    "TeamID" VARCHAR(15),
    "TeamTag" VARCHAR(50),
    "Country" VARCHAR(10),
    "Sponsor" VARCHAR(15)
);