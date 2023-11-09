CREATE TABLE IF NOT EXISTS "user_data" (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(80) NOT NULL,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO "user_data" (name)
VALUES ('Alan'),
    ('Cyan'),
    ('Dora'),
    ('Feynman'),
    ('Hanna'),
    ('Jacky'),
    ('Jason'),
    ('Kevin'),
    ('Nicky'),
    ('Valen'),
    ('Zoey');