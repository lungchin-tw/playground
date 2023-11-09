CREATE TABLE IF NOT EXISTS "class" (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(16) NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO "class" (title)
VALUES ('Sorcerer'),
    ('Barbarian'),
    ('Rogue'),
    ('Necromancer'),
    ('Druid');