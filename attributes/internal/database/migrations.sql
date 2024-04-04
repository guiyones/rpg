CREATE TABLE IF NOT EXISTS attributes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    attribute_type TEXT NOT NULL,
    description TEXT NOT NULL,
    value INTEGER NOT NULL
)