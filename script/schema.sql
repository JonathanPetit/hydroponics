SET timezone = 'UTC';

CREATE TABLE temperature (
    id serial PRIMARY KEY,
    value decimal,
    savedAt timestamp
);
