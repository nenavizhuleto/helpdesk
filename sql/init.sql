CREATE TABLE IF NOT EXISTS issues (
    id TEXT NOT NULL,
    client TEXT NOT NULL,
    company TEXT NOT NULL,
    department TEXT NOT NULL,
    name TEXT NOT NULL,
    phonenumber TEXT NOT NULL,
    innernumber TEXT NOT NULL,
    description TEXT NOT NULL,
    status TEXT NOT NULL,
    PRIMARY KEY(id, client)
)
